package gorilla

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewServer() {
	s := new(server)
	s.router = mux.NewRouter()
	AddListener(s)

}

type server struct {
	srv    *http.Server
	router *mux.Router
}

func (s *server) OnEvent(eventChannel <-chan Event) {

	for evt := range eventChannel {
		go evt.Execute(s)
	}

}
