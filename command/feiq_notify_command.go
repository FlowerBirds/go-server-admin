package command

import (
	"github.com/FlowerBirds/go-server-admin/base"
	"github.com/FlowerBirds/go-server-admin/config"
	event2 "github.com/FlowerBirds/go-server-admin/event"
	mess2 "github.com/FlowerBirds/go-server-admin/mess"
	"github.com/FlowerBirds/go-server-admin/util"
	"log"
	"strings"
	"time"
)

type FeiQNotifyCommand struct {
	*BaseCommand
	config     *config.UdpConfig
	sshClients map[string]*SshClient
}

type SshClient struct {
	Ip       string
	LastTime int64
}

func NewFeiQNotifyCommand(config *config.UdpConfig) *FeiQNotifyCommand {
	return &FeiQNotifyCommand{
		BaseCommand: &BaseCommand{
			OpType: base.FEIQ_NOTIFY_COMMAND,
		},
		config:     config,
		sshClients: make(map[string]*SshClient),
	}
}

func (c *FeiQNotifyCommand) OnRequest(event base.IEvent) {
	feiqEvent, ok := interface{}(event).(*event2.FeiQEvent)
	if ok {
		if util.EnableDebug() {
			log.Println("feiq notify： " + feiqEvent.Message.SshCommandLine)
		}
		clientMessage := feiqEvent.Message
		ip := clientMessage.SshClientIp
		now := time.Now().UnixMilli()
		if client, ok := c.sshClients[ip]; ok {
			if now-client.LastTime > 1000*3600 {
				mess := mess2.NewFeiQMessage("使用服务器期间，请谨慎操作，长时间未使用请及时断开", clientMessage.Ip, now)
				evt := event2.NewSendEvent()
				evt.SetToIp(ip)
				evt.Port = 2425
				evt.SetMessageOri(mess)
				evt.Gbk = true
				c.eventBus.PushEvent(evt)
			}
			client.LastTime = now
		} else {
			client := &SshClient{
				Ip:       ip,
				LastTime: now,
			}
			c.sshClients[ip] = client
			// 第一次登录执行命令，提示欢迎消息
			mess := mess2.NewFeiQMessage("欢迎登录服务器，请谨慎操作，切勿随意删除文件", clientMessage.Ip, now)
			evt := event2.NewSendEvent()
			evt.MessageOri = mess
			evt.SetToIp(ip)
			evt.Port = 2425
			evt.Gbk = true
			c.eventBus.PushEvent(evt)
		}
		// 检查执行命令，如果发现到rm删除命令，则立即通知
		if strings.Index(strings.ToLower(clientMessage.SshCommandLine), "rm ") >= 0 {
			mess := mess2.NewFeiQMessage("系统检测到您执行了删除命令，请谨慎操作："+clientMessage.SshCommandLine, clientMessage.Ip, now)
			evt := event2.NewSendEvent()
			evt.MessageOri = mess
			evt.SetToIp(ip)
			evt.Port = 2425
			evt.Gbk = true
			c.eventBus.PushEvent(evt)
		}
	}
}

func (c *FeiQNotifyCommand) OnResponse(event base.IEvent) {

}

func (c *FeiQNotifyCommand) AddToListener(bus *base.LiveEventBus) {
	c.eventBus = bus
	c.eventBus.AddListener(c)
}

func (c *FeiQNotifyCommand) OnEvent(event base.IEvent) {
	Trigger(c, event)
}
