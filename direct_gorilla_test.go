package gorilla

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestGorillaDirect(t *testing.T) {
	go listenAndServe()
	time.Sleep(1 * time.Second)
	resp, err := http.Get("http://localhost:8080/ping")

	log.Println(resp)
	log.Println(err)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

//Test gorilla round trip without bus mechanics to ensure dependencies are set up

func listenAndServe() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler)
	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listen on ")
	log.Fatal(srv.ListenAndServe())
}
