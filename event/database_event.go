package event

import "github.com/FlowerBirds/go-server-admin/base"

type DatabaseEvent struct {
	*base.Event
}

func NewDatabaseEvent() *DatabaseEvent {
	return &DatabaseEvent{
		Event: &base.Event{
			OpType: base.DATABASE_CRUD_COMMAND,
		},
	}
}
