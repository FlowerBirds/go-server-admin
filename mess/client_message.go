package mess

import (
	"errors"
	"fmt"
	"github.com/FlowerBirds/go-server-admin/base"
	"github.com/FlowerBirds/go-server-admin/util"
	"log"
	"strings"
)

/*
Udp Client data format:
	IP#TIME#USER#COMMANDLINE

Example:
	191.168.1.111#2022-07-01 12:11:21#root#rm -rf *

*/
type ClientMessage struct {
	*base.Message
	Mess           string
	SshClientIp    string
	SshCommandLine string
	SshOperateTime string
	SshUser        string
}

func (m *ClientMessage) Parse(instr string) (string, error) {
	str, err := m.Message.Parse(instr)
	if err != nil {
		return "", err
	}
	if len(str) == 0 {
		return "", errors.New("not find more info")
	}
	m.Mess = str
	if util.EnableDebug() {
		log.Println(str)
	}
	err = errors.New("unuseful client message")
	p1 := strings.Index(m.Mess, "#")
	if p1 <= 0 {
		return "", err
	}
	m.SshClientIp = m.Mess[0:p1]
	p2 := strings.Index(m.Mess[p1+1:], "#")
	if p2 <= 0 {
		return "", err
	}
	m.SshOperateTime = m.Mess[p1+1 : p1+1+p2]
	p3 := strings.Index(m.Mess[p1+1+p2+1:], "#")
	if p3 <= 0 {
		return "", err
	}
	m.SshUser = m.Mess[p1+1+p2+1 : p1+1+p2+1+p3]
	if len(m.Mess) <= p1+1+p2+1+p3+1 {
		return "", err
	}
	m.SshCommandLine = m.Mess[p1+1+p2+1+p3+1:]

	return "", nil
}

func (m *ClientMessage) Build() string {
	return fmt.Sprintf("%s%s#%s#%s#%s", m.Message.Build(), m.SshClientIp, m.SshOperateTime, m.SshUser, m.SshCommandLine)
}

func NewClientMessage() *ClientMessage {
	return &ClientMessage{
		Message: &base.Message{
			OpType: base.RECEIVE_CLIENT_COMMAND,
		},
	}
}
