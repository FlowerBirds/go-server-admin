package config

import (
	"container/list"
	"github.com/FlowerBirds/go-server-admin/base"
	"github.com/FlowerBirds/go-server-admin/util"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

const DEFAULT_UDP_PORT = 2345
const DEFAULT_HTTP_PORT = 2346
const DATABASE_FILE_NAME = "servers.db"

type UdpConfig struct {
	Port     int
	IsServer bool
	Servers  map[string]int64
	Clients  map[string]base.MachineInfo
	DataDir  string
	HttpPort int
}

func (config *UdpConfig) UpdateServer(ip string) {
	now := util.GetNowUnix()
	config.Servers[ip] = now
	// TODO 删除旧的服务器IP
}

func (config *UdpConfig) GetServers() *list.List {
	servers := list.New()
	for k, _ := range config.Servers {
		servers.PushBack(k)
	}
	return servers
}

func (config *UdpConfig) GetDataDir() string {
	if _, err := os.Stat(config.DataDir); os.IsNotExist(err) {
		os.Mkdir(config.DataDir, 0777)
	}
	return config.DataDir
}

func (config *UdpConfig) InitDatabase() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	dir := config.GetDataDir()
	file := filepath.Join(dir, DATABASE_FILE_NAME)
	orm.RegisterDataBase("default", "sqlite3", file)
	orm.RunSyncdb("default", false, true)

	config.Clients = make(map[string]base.MachineInfo)
	config.Servers = make(map[string]int64)
}

func (config *UdpConfig) UpdateClient(ip string, m base.MachineInfo) {
	config.Clients[ip] = m
}

func (config *UdpConfig) GetClients() map[string]base.MachineInfo {
	/*var clients []string
	for k, _ := range config.Clients {
		clients = append(clients, k)
	}*/
	return config.Clients
}
