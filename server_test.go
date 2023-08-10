package gorilla

import (
	bus "github.com/hexades/hexabus"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)

var s *server

func TestGorillaSuite(t *testing.T) {
	startServer(t)
	configureHandler(t)
	time.Sleep(2 * time.Second)

	//s.router.HandleFunc("/ping", PingHandler)

	resp, err := http.Get("http://localhost:8080/ping")

	log.Println(resp)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	log.Println(resp)
}

func startServer(t *testing.T) {
	log.Println("Start server...")
	s = NewServer()
	assert.NotNil(t, s)
	log.Println("Send server configuration information...")
	evt := NewSendEvent(ServerStart("localhost:8080", 15, 15))
	bus.Get().SendServerEvent(evt)

	log.Println("Done with start up.")
}

func configureHandler(t *testing.T) {
	evt := NewSendEvent(HandlerFunc("/ping", PingHandler))
	bus.Get().SendServerEvent(evt)

}
