package gorilla

import (
	"github.com/gorilla/mux"
	bus "github.com/hexades/hexabus"
	"net/http"
)

func NewServer() *server {
	s := new(server)
	s.router = mux.NewRouter()
	bus.AddServerListener(s)
	return s
}

type server struct {
	srv    *http.Server
	router *mux.Router
}

func (s *server) OnServerEvent(eventChannel <-chan bus.ServerEvent) {

	for serverEvent := range eventChannel {
		switch event := serverEvent.(type) {
		case GorillaEvent:
			go event.Execute(s)
		}
	}
}
