package event

import (
	"github.com/FlowerBirds/go-server-admin/base"
	"github.com/FlowerBirds/go-server-admin/mess"
)

type FeiQEvent struct {
	*base.Event
	Message *mess.ClientMessage
}

func NewFeiQEvent(message *mess.ClientMessage) *FeiQEvent {
	return &FeiQEvent{
		Event: &base.Event{
			OpType: base.FEIQ_NOTIFY_COMMAND,
		},
		Message: message,
	}
}
