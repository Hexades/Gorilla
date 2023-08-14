package hgorilla

type Event interface {
	Execute(server *server)
	Send(Response)
	Receive() Response
}
type EventModel struct {
	responseChannel chan Response
	executable      Executable
}

func NewEvent(executable Executable) Event {
	return &EventModel{executable: executable}
}

func (e *EventModel) Execute(server *server) {
	e.executable(server)
}

// TODO context time out on channel?
func (e *EventModel) getChannel() chan Response {
	if e.responseChannel == nil {
		e.responseChannel = make(chan Response, 1)
	}
	return e.responseChannel
}

func (e *EventModel) Send(val Response) {
	e.getChannel() <- val
}

func (e *EventModel) Receive() Response {
	return <-e.getChannel()
}
func newResponse(value any, err error) *Response {
	return &Response{value, err}
}

type Response struct {
	Value any
	Err   error
}
