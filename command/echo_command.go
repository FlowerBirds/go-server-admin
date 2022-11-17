package command

import (
	"github.com/FlowerBirds/go-server-admin/base"
)

type EchoCommand struct {
	*BaseCommand
}

func NewEchoCommand() *EchoCommand {
	return &EchoCommand{
		BaseCommand: &BaseCommand{
			OpType: base.ECHO_COMMAND,
		},
	}
}

func (c *EchoCommand) OnRequest(event base.IEvent) {

}

func (c *EchoCommand) OnResponse(event base.IEvent) {

}

func (c *EchoCommand) AddToListener(bus *base.LiveEventBus) {
	c.eventBus = bus
	c.eventBus.AddListener(c)
}
