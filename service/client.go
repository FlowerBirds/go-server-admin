package service

import (
	"github.com/FlowerBirds/go-server-admin/mess"
	"github.com/FlowerBirds/go-server-admin/util"
	"log"
	"strconv"
	"strings"
)

type UdpClient struct {
	ToIp   string
	ToPort int
	Data   string
}

func (client *UdpClient) Start() {
	if len(client.Data) == 0 {
		log.Println("not find any data to send.")
		return
	}
	if util.EnableDebug() {
		log.Println("send to " + client.ToIp + ":" + strconv.Itoa(client.ToPort))
		log.Println(client.Data)
	}
	data := client.Data
	if strings.Index(client.Data, "003:") == 0 {
		m := mess.NewClientMessage()
		m.Parse(client.Data)
		if m.Ip == "127.0.0.1" {
			m.Ip = util.GetIp()
		}
		m.SshOperateTime = strconv.FormatInt(util.GetNowUnix(), 10)
		data = m.Build()
	}
	if util.EnableDebug() {
		log.Println("client send: " + client.Data)
	}

	util.SendUdpPoint(client.ToIp, client.ToPort, data, false)
}
