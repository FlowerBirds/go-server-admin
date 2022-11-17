package base

type IEvent interface {
	GetOpType() int
	SetOpType(op int)
	GetInOut() int
	SetInOut(inOut int)
	GetFromIp() string
	SetFromIp(fromIp string)
	GetToIP() string
	SetToIp(ip string)
	GetMessageStr() string
	SetMessageStr(str string)
	GetMessageOri() IMessage
}
