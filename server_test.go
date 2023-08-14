package hgorilla

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func TestGorillaSuite(t *testing.T) {
	startServer(t)
	configureHandler(t)
	//time.Sleep(10 * time.Second)

	//s.router.HandleFunc("/ping", PingHandler)

	resp, err := http.Get("http://localhost:8080/ping")
	log.Println(resp.StatusCode)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func startServer(t *testing.T) {
	log.Println("Start server...")
	NewServer()
	log.Println("Send server configuration information...")
	evt := NewEvent(ServerStart("localhost:8080", 15, 15))
	SendEvent(evt)
	log.Println("Done with start up.")
}

func configureHandler(t *testing.T) {
	evt := NewEvent(HandlerFunc("/ping", PingHandler))
	SendEvent(evt)

}
