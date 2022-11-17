package base

type IMessage interface {
	GetOpType() int
	SetOpType(op int)
	GetInOut() int
	SetInOut(inOut int)
	GetIp() string
	SetIp(ip string)
	Build() string
	Parse(mess string) (string, error)
}
