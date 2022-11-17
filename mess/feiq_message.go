package mess

import (
	"fmt"
	"github.com/FlowerBirds/go-server-admin/base"
)

type FeiQMessage struct {
	*base.Message
	Mess     string
	ServerIp string
	LastTime int64
}

func (m *FeiQMessage) Parse(instr string) (string, error) {
	m.Mess = instr
	return "", nil
}

func (m *FeiQMessage) Build() string {
	return fmt.Sprintf("1:%d:%s:%s:32:%s", m.LastTime, "Server Admin", m.ServerIp, m.Mess)
}

func NewFeiQMessage(mess string, ip string, lastTime int64) *FeiQMessage {
	return &FeiQMessage{
		Message: &base.Message{
			OpType: base.FEIQ_NOTIFY_COMMAND,
		},
		Mess:     mess,
		ServerIp: ip,
		LastTime: lastTime,
	}
}
