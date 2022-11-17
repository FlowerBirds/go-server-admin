package base

import (
	"errors"
	"fmt"
	"github.com/FlowerBirds/go-server-admin/util"
	"log"
	"strconv"
	"strings"
)

const REQUEST_MESSAGE = 0
const RESPONSE_MESSAGE = 1

type Message struct {
	Ip     string
	InOut  int
	OpType int
}

func (m *Message) Build() string {
	return fmt.Sprintf("%03d:%d:%s:", m.GetOpType(), m.GetInOut(), m.GetIp())
}

func (m *Message) Parse(mess string) (string, error) {
	if len(mess) < 0 {
		return "", errors.New("not find any message")
	}
	var err error
	m.OpType, err = strconv.Atoi(mess[0:3])
	if err != nil {
		return "", err
	}
	m.InOut, err = strconv.Atoi(mess[4:5])
	if err != nil {
		return "", err
	}
	i := strings.Index(mess[7:], ":")
	if i <= 0 {
		return "", errors.New("not find client ip info")
	}
	m.Ip = mess[6 : 7+i]
	if util.EnableDebug() {
		log.Printf("%s %s %s", mess[0:3], mess[4:5], mess[6:7+i])
	}
	if len(mess) < 7+i+1 {
		return "", errors.New("not find more info for other message")
	}

	return mess[7+i+1:], nil
}

func NewMessage(str string) (*Message, error) {
	message := &Message{}
	_, err := message.Parse(str)
	return message, err
}

func (m *Message) GetOpType() int {
	return m.OpType
}

func (m *Message) SetOpType(op int) {
	m.OpType = op
}

func (m *Message) GetInOut() int {
	return m.InOut
}

func (m *Message) SetInOut(inOut int) {
	m.InOut = inOut
}

func (m *Message) GetIp() string {
	return m.Ip
}

func (m *Message) SetIp(ip string) {
	m.Ip = ip
}
