package event

import (
	"github.com/FlowerBirds/go-server-admin/base"
)

type UpdateGetEvent struct {
	*base.Event
	ServerIp      string
	ServerVersion string
}

func NewUpdateGetEvent(ip string, version string) *UpdateGetEvent {
	return &UpdateGetEvent{
		Event: &base.Event{
			OpType: base.UPDATE_GET_COMMAND,
		},
		ServerIp:      ip,
		ServerVersion: version,
	}
}
