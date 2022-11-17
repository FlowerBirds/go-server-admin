package dao

import (
	"github.com/FlowerBirds/go-server-admin/model"
	"github.com/astaxie/beego/orm"
)

func ListClientMessages(currentPage int, pageSize int, currentClient string) []model.ClientMessageModel {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(model.ClientMessageModel))
	var messages []model.ClientMessageModel
	// qs.All(&messages)
	if len(currentClient) == 0 {
		qs.OrderBy("-serverTime").Limit(pageSize, pageSize*(currentPage-1)).All(&messages)
	} else {
		qs.Filter("clientIp", currentClient).OrderBy("-serverTime").Limit(pageSize, pageSize*(currentPage-1)).All(&messages)
	}
	return messages
}

func CountClientMessages(currentClient string) int64 {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(new(model.ClientMessageModel))
	var total int64 = 0
	if len(currentClient) == 0 {
		total, _ = qs.Count()
	} else {
		total, _ = qs.Filter("clientIp", currentClient).Count()
	}

	return total
}
