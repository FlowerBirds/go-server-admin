package base

type Task struct {
	event   IEvent
	command ICommand
}

func (t *Task) GetEvent() IEvent {
	return t.event
}

func (t *Task) GetCommand() ICommand {
	return t.command
}
