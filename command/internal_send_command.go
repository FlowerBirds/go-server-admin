package command

import (
	"github.com/FlowerBirds/go-server-admin/base"
	"github.com/FlowerBirds/go-server-admin/config"
	event2 "github.com/FlowerBirds/go-server-admin/event"
	"github.com/FlowerBirds/go-server-admin/util"
	"log"
)

type SendCommand struct {
	*BaseCommand
}

func NewSendCommand() *SendCommand {
	return &SendCommand{
		BaseCommand: &BaseCommand{
			OpType: base.INTERNAL_COMMAND,
		},
	}
}

func (c *SendCommand) OnRequest(evt base.IEvent) {
	evt.GetMessageOri().SetIp(util.GetIp())
	sendEvent, ok := interface{}(evt).(*event2.SendEvent)
	if ok {
		port := config.DEFAULT_UDP_PORT
		if sendEvent.Port > 0 {
			port = sendEvent.Port
		}
		if util.EnableDebug() {
			log.Printf("send to %s:%d\n", sendEvent.GetToIP(), port)
		}
		util.SendUdpPoint(sendEvent.GetToIP(), port, evt.GetMessageOri().Build(), sendEvent.Gbk)
	}
	_, ok = interface{}(evt).(*event2.SendAllEvent)
	if ok {
		util.SendMultiPointFrom1To3(config.DEFAULT_UDP_PORT, evt.GetMessageOri().Build())
	}
}

func (c *SendCommand) OnResponse(event base.IEvent) {

}

func (c *SendCommand) AddToListener(bus *base.LiveEventBus) {
	c.eventBus = bus
	c.eventBus.AddListener(c)
}

func (c *SendCommand) OnEvent(event base.IEvent) {
	Trigger(c, event)
}
