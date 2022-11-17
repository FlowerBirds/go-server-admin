package service

import (
	"encoding/json"
	"github.com/FlowerBirds/go-server-admin/base"
	"github.com/FlowerBirds/go-server-admin/config"
	"github.com/FlowerBirds/go-server-admin/event"
	"github.com/FlowerBirds/go-server-admin/mess"
	"github.com/FlowerBirds/go-server-admin/util"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"log"
	"time"
)

type EdgeMonitorService struct {
	EventBus base.LiveEventBus
	Config   *config.UdpConfig
}

func (service *EdgeMonitorService) Start() {
	log.Println("start edge monitor service")
	service.handler(15)

}

func (service *EdgeMonitorService) handler(sec uint) {
	starter := time.NewTimer(time.Second * time.Duration(sec))
	<-starter.C
	starter.Stop()

	m := service.getMachineInfo()
	b, err := json.Marshal(m)

	if err != nil {
		log.Println(err)
	} else {
		monitorEvent := event.NewMonitorEvent()
		message := mess.NewMonitorMessage()
		message.Machine = string(b)
		monitorEvent.SetMessageOri(message)
		service.EventBus.PushEvent(monitorEvent)
	}

	service.handler(15)
}

func (service EdgeMonitorService) getMachineInfo() base.MachineInfo {
	info, _ := host.Info()
	total, free, ava := service.getMemory()
	cores, f := service.getCpu()
	l1, l5, l15 := service.getCpuLoad()

	m := base.MachineInfo{
		Hostname: info.Hostname,
		OS:       info.OS,
		Mem: base.MemInfo{
			Total:     total,
			Free:      free,
			Available: ava,
		},
		Cpu: base.CpuInfo{
			Cores:  cores,
			Family: f,
			Load1:  l1,
			Load5:  l5,
			Load15: l15,
		},
		Disk: service.getDisk(),
		User: service.getUsers(),
		Time: util.GetNowUnix(),
	}
	return m
}

func (service *EdgeMonitorService) getMemory() (uint64, uint64, uint64) {
	v, _ := mem.VirtualMemory()
	return v.Total, v.Free, v.Available
}

func (service *EdgeMonitorService) getCpu() (int, string) {
	info, _ := cpu.Info()
	cores := len(info)
	family := "unknown"
	if cores > 0 {
		family = info[0].Family
	}
	return cores, family
}

func (service *EdgeMonitorService) getCpuLoad() (float64, float64, float64) {
	ld, _ := load.Avg()
	return ld.Load1, ld.Load5, ld.Load15
}

func (service *EdgeMonitorService) getDisk() []base.DiskInfo {
	partitions, _ := disk.Partitions(false)
	var disks []base.DiskInfo
	for _, v := range partitions {
		usage, _ := disk.Usage(v.Mountpoint)
		d := base.DiskInfo{
			Path:  v.Mountpoint,
			Total: usage.Total,
			Free:  usage.Free,
		}
		disks = append(disks, d)
	}
	return disks
}

func (service *EdgeMonitorService) getUsers() []base.UserInfo {
	var users []base.UserInfo
	us, _ := host.Users()
	for _, v := range us {
		u := base.UserInfo{
			Host:    v.Host,
			User:    v.User,
			Started: v.Started,
		}
		users = append(users, u)
	}
	return users
}
