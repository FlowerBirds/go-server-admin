package command

import (
	"github.com/FlowerBirds/go-server-admin/base"
	"github.com/FlowerBirds/go-server-admin/config"
	event2 "github.com/FlowerBirds/go-server-admin/event"
	"github.com/FlowerBirds/go-server-admin/mess"
	"github.com/FlowerBirds/go-server-admin/util"
	"log"
)

type GetServerCommand struct {
	*BaseCommand
	config *config.UdpConfig
}

func NewGetServerCommand(config *config.UdpConfig) *GetServerCommand {
	return &GetServerCommand{
		BaseCommand: &BaseCommand{
			OpType: base.GET_SERVER_COMMAND,
		},
		config: config,
	}
}

func (c *GetServerCommand) OnRequest(event base.IEvent) {
	if util.EnableDebug() {
		log.Println("receive from client: " + event.GetFromIp() + " | " + event.GetMessageStr())
	}
	getServerMessage := mess.NewGetServerMessage()
	getServerMessage.SetInOut(base.RESPONSE_MESSAGE)
	getServerMessage.ServerIp = util.GetIp()
	// 判断客户端请求的参数信息，是否已经携带了版本信息，如果携带，则将服务端的版本信息传回
	clientMessage := mess.NewGetServerMessage()
	clientMessage.Parse(event.GetMessageStr())
	if len(clientMessage.Version) > 0 {
		getServerMessage.Version = c.config.Version
	}
	sendEvent := event2.NewSendEvent()
	sendEvent.MessageOri = getServerMessage
	sendEvent.SetToIp(event.GetFromIp())

	c.eventBus.PushEvent(sendEvent)
}

func (c *GetServerCommand) OnResponse(event base.IEvent) {
	if util.EnableDebug() {
		log.Println("receive from server " + event.GetFromIp())
	}
	serverMessage := mess.NewGetServerMessage()
	serverMessage.Parse(event.GetMessageStr())
	c.config.UpdateServer(serverMessage.ServerIp)
	if util.EnableDebug() && len(serverMessage.Version) > 0 {
		log.Println("server version: " + serverMessage.Version)
	}
	// 如果服务端传回了其版本号，需要进行对比，然后再决定是否进行更新
	if len(serverMessage.Version) > 0 && util.CompareVersion(c.config.Version, serverMessage.Version) < 0 {
		// 此时需要更新本地程序

	}

}

func (c *GetServerCommand) OnEvent(event base.IEvent) {
	Trigger(c, event)
}

func (c *GetServerCommand) AddToListener(bus *base.LiveEventBus) {
	c.eventBus = bus
	c.eventBus.AddListener(c)
}
