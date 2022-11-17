package event

import (
	"github.com/FlowerBirds/go-server-admin/base"
)

type SendEvent struct {
	*base.Event
	Port int
	Gbk  bool
}

func NewSendEvent() *SendEvent {
	return &SendEvent{
		Event: &base.Event{
			OpType: base.INTERNAL_COMMAND,
		},
	}
}
