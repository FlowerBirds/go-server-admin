package service

import (
	"github.com/FlowerBirds/go-server-admin/base"
	"github.com/FlowerBirds/go-server-admin/config"
	"github.com/FlowerBirds/go-server-admin/event"
	"github.com/FlowerBirds/go-server-admin/mess"
	"github.com/FlowerBirds/go-server-admin/util"
	"log"
	"time"
)

type GetServerService struct {
	EventBus base.LiveEventBus
	Config   *config.UdpConfig
}

func (service *GetServerService) Start() {
	log.Println("start get server service")
	service.handler(5)
}

func (service *GetServerService) handler(sec uint) {
	starter := time.NewTimer(time.Second * time.Duration(sec))
	<-starter.C
	starter.Stop()
	service.getServerInfo()
	// 当发现的服务器列表大于1，即除过自身外的其它服务，则降低获取服务列表的频率
	if service.Config.GetServers().Len() > 1 {
		service.handler(60)
	} else {
		service.handler(15)
	}
}

func (service *GetServerService) getServerInfo() {
	if util.EnableDebug() {
		log.Println("Try to get server info.")
	}
	getServerMessage := mess.NewGetServerMessage()
	// 添加客户端的版本号
	getServerMessage.Version = service.Config.Version
	// getServerMessage.ServerIp = util.GetIp()
	sendAllEvent := event.NewSendAllEvent()
	sendAllEvent.MessageOri = getServerMessage

	service.EventBus.PushEvent(sendAllEvent)

}
