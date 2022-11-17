package mess

import (
	"github.com/FlowerBirds/go-server-admin/base"
)

type EchoMessage struct {
	*base.Message
	Mess string
}

func NewEchoMessage() *EchoMessage {
	return &EchoMessage{
		Message: &base.Message{
			OpType: base.ECHO_COMMAND,
		},
	}
}

func (m *EchoMessage) Parse(instr string) (string, error) {
	str, err := m.Message.Parse(instr)
	if err != nil {
		return "", err
	}
	if len(str) > 0 {
		m.Mess = str
	}
	return "", err
}

func (m *EchoMessage) Build() string {
	return m.Message.Build() + ":" + m.Mess
}
