package mess

import (
	"errors"
	"github.com/FlowerBirds/go-server-admin/base"
)

type GetServerMessage struct {
	*base.Message
	ServerIp string
}

func NewGetServerMessage() *GetServerMessage {
	return &GetServerMessage{
		Message: &base.Message{
			OpType: base.GET_SERVER_COMMAND,
		},
	}
}

func (m *GetServerMessage) Parse(mess string) (string, error) {
	str, err := m.Message.Parse(mess)
	if err != nil {
		return "", err
	}
	if len(str) <= 1 {
		return "", errors.New("not find server ip")
	}
	m.ServerIp = str
	return "", nil
}

func (m *GetServerMessage) Build() string {
	return m.Message.Build() + m.ServerIp
}
