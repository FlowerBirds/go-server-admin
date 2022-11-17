package command

import (
	"github.com/FlowerBirds/go-server-admin/base"
	"github.com/FlowerBirds/go-server-admin/config"
)

type UpdateGetCommand struct {
	*BaseCommand
}

func NewUpdateGetCommand(config *config.UdpConfig) *UpdateGetCommand {
	return &UpdateGetCommand{
		BaseCommand: &BaseCommand{
			OpType: base.UPDATE_GET_COMMAND,
		},
	}
}

func (c *UpdateGetCommand) OnRequest(event base.IEvent) {
	// 接收的版本不一致信息后，记录下来等待更新，更新过程中不再处理其它请求
	// 一方面：本地需要准备好文件接收
	// 另外一方面：需要告诉服务器端，客户端已经准备好，随时接收文件，消息中携带本地接听的端口

}

func (c *UpdateGetCommand) OnResponse(event base.IEvent) {

}

func (c *UpdateGetCommand) AddToListener(bus *base.LiveEventBus) {
	c.eventBus = bus
	c.eventBus.AddListener(c)
}

func (c *UpdateGetCommand) OnEvent(event base.IEvent) {
	Trigger(c, event)
}
