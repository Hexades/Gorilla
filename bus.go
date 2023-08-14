package hgorilla

var bus = gorillabus{}

type gorillabus struct {
	listenerChannels []chan Event
}

func AddListener(eventListener EventListener) {
	listenerChannel := make(chan Event, 10)
	bus.listenerChannels = append(bus.listenerChannels, listenerChannel)
	go eventListener.onEvent(listenerChannel)
}

func SendEvent(event Event) {
	for _, channel := range bus.listenerChannels {
		channel <- event
	}
}

type EventListener interface {
	onEvent(eventChannel <-chan Event)
}
