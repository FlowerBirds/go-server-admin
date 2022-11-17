package command

import (
	"github.com/FlowerBirds/go-server-admin/base"
	"github.com/FlowerBirds/go-server-admin/mess"
	"github.com/FlowerBirds/go-server-admin/model"
	"github.com/FlowerBirds/go-server-admin/util"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)

type DataBaseCommand struct {
	*BaseCommand
	orm orm.Ormer
}

func NewDataBaseCommand() *DataBaseCommand {
	o := orm.NewOrm()
	o.Using("default")
	return &DataBaseCommand{
		BaseCommand: &BaseCommand{
			OpType: base.DATABASE_CRUD_COMMAND,
		},
		orm: o,
	}
}

func (c *DataBaseCommand) OnRequest(event base.IEvent) {
	oriMess := event.GetMessageOri()

	if util.IsInstanceOf(oriMess, &mess.ClientMessage{}) {
		clientMess, _ := interface{}(oriMess).(*mess.ClientMessage)
		model := new(model.ClientMessageModel)
		model.SshUser = clientMess.SshUser
		model.SshClientIp = clientMess.SshClientIp
		model.ServerTime = util.GetNowUnix()
		model.ClientIp = clientMess.Ip
		t, _ := strconv.ParseInt(clientMess.SshOperateTime, 10, 64)
		model.SshOperateTime = t
		if len(clientMess.SshCommandLine) >= 255 {
			model.SshCommandLine = clientMess.SshCommandLine[:255]
		} else {
			model.SshCommandLine = clientMess.SshCommandLine
		}
		// 保存到数据库
		id, err := c.orm.Insert(model)
		if err != nil {
			log.Println(err)
		} else if util.EnableDebug() {
			log.Printf("insert: %d\n", id)
		}
	} else if util.IsInstanceOf(oriMess, &mess.MonitorMessage{}) {
		monitorMess, _ := interface{}(oriMess).(*mess.MonitorMessage)
		log.Println("save table: " + monitorMess.Machine)
	}

}

func (c *DataBaseCommand) OnResponse(event base.IEvent) {

}

func (c *DataBaseCommand) AddToListener(bus *base.LiveEventBus) {
	c.eventBus = bus
	c.eventBus.AddListener(c)
}

func (c *DataBaseCommand) OnEvent(event base.IEvent) {
	Trigger(c, event)
}
