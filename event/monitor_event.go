package event

import "github.com/FlowerBirds/go-server-admin/base"

type MonitorEvent struct {
	*base.Event
}

func NewMonitorEvent() *MonitorEvent {
	return &MonitorEvent{
		Event: &base.Event{
			OpType: base.EDGE_MONITOR_COMMAND,
		},
	}
}
