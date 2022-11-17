package command

import (
	"github.com/FlowerBirds/go-server-admin/base"
	"log"
)

type BaseCommand struct {
	eventBus *base.LiveEventBus
	OpType   int
}

func (c *BaseCommand) AddToListener(eventBus *base.LiveEventBus) {
	c.eventBus = eventBus
	c.eventBus.AddListener(c)
}

func (c *BaseCommand) OnEvent(event base.IEvent) {
	Trigger(c, event)
}

func (c *BaseCommand) GetType() int {
	return c.OpType
}

func (c *BaseCommand) dealMessage(event base.IEvent) {
	if event.GetInOut() == base.REQUEST_MESSAGE {
		c.OnRequest(event)
	} else if event.GetInOut() == base.RESPONSE_MESSAGE {
		c.OnResponse(event)
	}
}

func (c *BaseCommand) PushEvent(event base.IEvent) {
	c.eventBus.PushEvent(event)
}

func (c *BaseCommand) OnRequest(event base.IEvent) {
	log.Println("onRequest nothing todo")
}

func (c *BaseCommand) OnResponse(event base.IEvent) {
	log.Println("onResponse nothing todo")
}

func Trigger(c base.ICommand, event base.IEvent) {
	if event.GetInOut() == base.REQUEST_MESSAGE {
		c.OnRequest(event)
	} else if event.GetInOut() == base.RESPONSE_MESSAGE {
		c.OnResponse(event)
	}
}
