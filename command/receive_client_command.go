package command

import (
	"github.com/FlowerBirds/go-server-admin/base"
	"github.com/FlowerBirds/go-server-admin/config"
	event2 "github.com/FlowerBirds/go-server-admin/event"
	"github.com/FlowerBirds/go-server-admin/mess"
	"github.com/FlowerBirds/go-server-admin/util"
	"log"
)

type ReceiveClientCommand struct {
	*BaseCommand
	config *config.UdpConfig
}

func NewReceiveClientCommand(config *config.UdpConfig) *ReceiveClientCommand {
	return &ReceiveClientCommand{
		BaseCommand: &BaseCommand{
			OpType: base.RECEIVE_CLIENT_COMMAND,
		},
		config: config,
	}
}

/**
  receive data from udp client to server admin client, then send to server admin server edg.
*/
func (c *ReceiveClientCommand) OnRequest(event base.IEvent) {
	// 获取到客户端发送过来的数据，需要发送给服务端进行存储
	// 获取服务器列表，组装数据，发送
	clientMessage := mess.NewClientMessage()
	clientMessage.Parse(event.GetMessageStr())
	clientMessage.SetInOut(base.RESPONSE_MESSAGE)
	if util.EnableDebug() {
		log.Println("receive client: " + event.GetMessageStr())
	}
	for i := c.config.GetServers().Front(); i != nil; i = i.Next() {
		ip := i.Value.(string)
		sendEvent := event2.NewSendEvent()
		sendEvent.SetToIp(ip)
		sendEvent.SetMessageOri(clientMessage)
		c.eventBus.PushEvent(sendEvent)
	}
	// 根据收到的客户端操作信息，判断是否进行了危险操作，否则通过飞秋发出提示信息
	feiqEvent := event2.NewFeiQEvent(clientMessage)
	c.eventBus.PushEvent(feiqEvent)

}

/**
receive from server admin client not udp client
*/
func (c *ReceiveClientCommand) OnResponse(event base.IEvent) {
	clientMessage := mess.NewClientMessage()
	clientMessage.Parse(event.GetMessageStr())
	// 将接受到客户端服务器的消息，并记录客户端服务器的IP，保存到数据库中
	clientMessage.Ip = event.GetFromIp()
	// 此时需要将客户端收集的操作日志信息给存储起来
	databaseEvent := event2.NewDatabaseEvent()
	databaseEvent.SetMessageOri(clientMessage)

	c.eventBus.PushEvent(databaseEvent)
}

func (c *ReceiveClientCommand) AddToListener(bus *base.LiveEventBus) {
	c.eventBus = bus
	c.eventBus.AddListener(c)
}

func (c *ReceiveClientCommand) OnEvent(event base.IEvent) {
	Trigger(c, event)
}
