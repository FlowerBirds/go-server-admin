package base

type ICommand interface {
	AddToListener(bus *LiveEventBus)

	OnEvent(event IEvent)

	GetType() int

	OnRequest(event IEvent)

	OnResponse(event IEvent)
}
