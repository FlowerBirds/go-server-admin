package model

import "github.com/astaxie/beego/orm"

type ClientMessageModel struct {
	Id             int64  `json:"id"`
	SshClientIp    string `json:"sshClientIp"`
	SshCommandLine string `orm:"size(2000);" json:"sshCommandLine"`
	SshOperateTime int64  `json:"sshOperateTime"`
	SshUser        string `orm:"size(100);" json:"sshUser"`
	ServerTime     int64  `json:"serverTime"`
	ClientIp       string `json:"clientIp,omitempty"`
}

type ClientMessageVO struct {
	Total       int64                `json:"total"`
	Data        []ClientMessageModel `json:"data"`
	CurrentPage int                  `json:"currentPage"`
	PageSize    int                  `json:"pageSize"`
}

func init() {
	orm.RegisterModel(new(ClientMessageModel))
}
