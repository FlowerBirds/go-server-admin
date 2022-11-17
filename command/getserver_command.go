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
		log.Println("receive from client: " + event.GetFromIp())
	}
	getServerMessage := mess.NewGetServerMessage()
	getServerMessage.SetInOut(base.RESPONSE_MESSAGE)
	getServerMessage.ServerIp = util.GetIp()
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

}

func (c *GetServerCommand) OnEvent(event base.IEvent) {
	Trigger(c, event)
}

func (c *GetServerCommand) AddToListener(bus *base.LiveEventBus) {
	c.eventBus = bus
	c.eventBus.AddListener(c)
}
