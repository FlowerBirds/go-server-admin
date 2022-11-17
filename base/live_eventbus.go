package base

import (
	"errors"
	"github.com/FlowerBirds/go-server-admin/util"
	"log"
	"reflect"
	"strconv"
)

type LiveEventBus struct {
	listeners map[int]ICommand
	tasks     chan Task
}

func NewEventBus() *LiveEventBus {
	listeners := make(map[int]ICommand)
	tasks := make(chan Task, 100)
	eventbus := &LiveEventBus{
		listeners: listeners,
		tasks:     tasks,
	}
	eventbus.start()
	return eventbus
}

func (eventbus *LiveEventBus) AddListener(command ICommand) error {
	if _, ok := eventbus.listeners[command.GetType()]; ok {
		err := errors.New("Command has exist: " + strconv.Itoa(command.GetType()))
		return err
	} else {
		eventbus.listeners[command.GetType()] = command
		log.Printf("add command: %s for %d\n", reflect.TypeOf(command).String(), command.GetType())
	}
	return nil
}

func (eventbus *LiveEventBus) PushEvent(event IEvent) {
	if command, ok := eventbus.listeners[event.GetOpType()]; ok {
		// 将任务放入管道中，等待被消费
		eventbus.tasks <- Task{
			event:   event,
			command: command,
		}
		if util.EnableDebug() {
			log.Println("push event: " + reflect.TypeOf(command).String())
		}
	} else {
		log.Println("unknown event: " + strconv.Itoa(event.GetOpType()))
	}
}

func (eventbus *LiveEventBus) start() {
	go eventbus.doTask()
}

func (eventbus *LiveEventBus) doTask() {
	for v := range eventbus.tasks {
		// log.Println("task with " + strconv.Itoa(v.GetEvent().GetOpType()))
		v.GetCommand().OnEvent(v.GetEvent())
	}
}
