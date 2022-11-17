package mess

import (
	"fmt"
	"github.com/FlowerBirds/go-server-admin/base"
)

type MonitorMessage struct {
	*base.Message
	Machine string
}

func NewMonitorMessage() *MonitorMessage {
	return &MonitorMessage{
		Message: &base.Message{
			OpType: base.EDGE_MONITOR_COMMAND,
		},
	}
}

func (m *MonitorMessage) Build() string {
	return fmt.Sprintf("%s%s", m.Message.Build(), m.Machine)
}

func (m *MonitorMessage) Parse(instr string) (string, error) {
	str, err := m.Message.Parse(instr)
	if err != nil {
		return "", err
	}
	if len(str) > 0 {
		m.Machine = str
	}
	return "", err
}
