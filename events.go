package gorilla

type GorillaEvent interface {
	Execute(server *server)
}

func NewSendEvent(sendFunc AsyncFunc) *SendEvent {
	return &SendEvent{asyncFunc: sendFunc}
}

type RequestResponseEvent struct {
	GorillaEvent
	serverFunc SynchronousFunc
}

type SendEvent struct {
	GorillaEvent
	asyncFunc AsyncFunc
}

func (e *SendEvent) Execute(server *server) {
	e.asyncFunc(server)

}
