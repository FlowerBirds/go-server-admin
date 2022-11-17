package command

import (
	"encoding/json"
	"github.com/FlowerBirds/go-server-admin/base"
	"github.com/FlowerBirds/go-server-admin/config"
	event2 "github.com/FlowerBirds/go-server-admin/event"
	"github.com/FlowerBirds/go-server-admin/mess"
	"github.com/FlowerBirds/go-server-admin/util"
	"log"
	"strconv"
)

type EdgeMonitorCommand struct {
	*BaseCommand
	config *config.UdpConfig
}

func NewEdgeMonitorCommand(config *config.UdpConfig) *EdgeMonitorCommand {
	return &EdgeMonitorCommand{
		BaseCommand: &BaseCommand{
			OpType: base.EDGE_MONITOR_COMMAND,
		},
		config: config,
	}
}

func (c *EdgeMonitorCommand) OnRequest(event base.IEvent) {
	monitorMessage := mess.NewMonitorMessage()
	oldMess := interface{}(event.GetMessageOri()).(*mess.MonitorMessage)
	monitorMessage.SetInOut(base.RESPONSE_MESSAGE)
	monitorMessage.Machine = oldMess.Machine
	if util.EnableDebug() {
		log.Println("monitor OnRequest: " + oldMess.Machine)
	}
	for i := c.config.GetServers().Front(); i != nil; i = i.Next() {
		ip := i.Value.(string)
		sendEvent := event2.NewSendEvent()
		sendEvent.SetToIp(ip)
		sendEvent.SetMessageOri(monitorMessage)
		c.eventBus.PushEvent(sendEvent)
	}
}

func (c *EdgeMonitorCommand) OnResponse(event base.IEvent) {
	monitorMessage := mess.NewMonitorMessage()
	monitorMessage.Parse(event.GetMessageStr())
	monitorMessage.Ip = event.GetFromIp()

	/**
	databaseEvent := event2.NewDatabaseEvent()
	databaseEvent.SetMessageOri(monitorMessage)

	c.eventBus.PushEvent(databaseEvent)

	*/

	m := base.MachineInfo{}
	json.Unmarshal([]byte(monitorMessage.Machine), &m)
	if util.EnableDebug() {
		log.Println(m.Hostname + " " + strconv.FormatInt(m.Time, 10))
	}

	// 更新到客户端列表中
	c.config.UpdateClient(monitorMessage.Ip, m)

}

func (c *EdgeMonitorCommand) AddToListener(bus *base.LiveEventBus) {
	c.eventBus = bus
	c.eventBus.AddListener(c)
}

func (c *EdgeMonitorCommand) OnEvent(event base.IEvent) {
	Trigger(c, event)
}
