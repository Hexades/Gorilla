package gorilla

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Executable = func(*server) Response

var PingHandler = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("Ping")
	w.Write([]byte("Pong!"))
}

var ServerStart = func(host string, readTimeout, writeTimeout int64) Executable {
	return func(server *server) Response {

		server.router = mux.NewRouter()
		//TODO The time out multiplication time seconds works with const value but not passed in integers. TBD.
		server.srv = &http.Server{
			Handler:      server.router,
			Addr:         host,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		go server.srv.ListenAndServe()
		return Response{"Listening on address: " + server.srv.Addr, nil}

	}
}
var HandlerFunc = func(path string, handler func(w http.ResponseWriter, r *http.Request)) Executable {
	return func(server *server) Response {
		log.Println("Setting server: ", server)
		server.router.HandleFunc(path, handler)
		return Response{"OK", nil}
	}
}
