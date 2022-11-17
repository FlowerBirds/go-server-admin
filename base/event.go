package base

type Event struct {
	OpType     int
	InOut      int
	FromIp     string
	ToIP       string
	MessageStr string
	MessageOri IMessage
}

func (e *Event) GetOpType() int {
	return e.OpType
}

func (e *Event) SetOpType(op int) {
	e.OpType = op
}

func (e *Event) GetInOut() int {
	return e.InOut
}

func (e *Event) SetInOut(inOut int) {
	e.InOut = inOut
}

func (e *Event) GetFromIp() string {
	return e.FromIp
}

func (e *Event) SetFromIp(fromIp string) {
	e.FromIp = fromIp
}

func (e *Event) GetToIP() string {
	return e.ToIP
}

func (e *Event) SetToIp(ip string) {
	e.ToIP = ip
}

func (e *Event) GetMessageStr() string {
	return e.MessageStr
}

func (e *Event) SetMessageStr(str string) {
	e.MessageStr = str
}

func (e *Event) GetMessageOri() IMessage {
	return e.MessageOri
}

func (e *Event) SetMessageOri(messageOri IMessage) {
	e.MessageOri = messageOri
}
