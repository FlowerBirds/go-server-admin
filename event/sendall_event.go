package event

type SendAllEvent struct {
	*SendEvent
}

func NewSendAllEvent() *SendAllEvent {
	return &SendAllEvent{
		SendEvent: NewSendEvent(),
	}
}
