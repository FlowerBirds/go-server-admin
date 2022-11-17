package mess

import (
	"errors"
	"github.com/FlowerBirds/go-server-admin/base"
	"strings"
)

type GetServerMessage struct {
	*base.Message
	ServerIp string
	Version  string
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
	if strings.Index(str, ":") > -1 {
		arr := strings.Split(str, ":")
		if len(arr) == 2 {
			m.ServerIp = arr[0]
			m.Version = arr[1]
		}
	} else {
		m.ServerIp = str
	}
	return "", nil
}

func (m *GetServerMessage) Build() string {
	extInfo := ""
	if len(m.Version) > 0 {
		extInfo = ":" + m.Version
	}
	return m.Message.Build() + m.ServerIp + extInfo
}
