package base

type ITask interface {
	GetEvent() IEvent
	GetCommand() ICommand
}
