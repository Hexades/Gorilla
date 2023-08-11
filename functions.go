package gorilla

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	bus "github.com/hexades/hexabus"
)

type SynchronousFunc = func(server *server) bus.Response
type AsyncFunc = func(server *server)

var PingHandler = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var ServerStart = func(host string, readTimeout, writeTimeout int64) AsyncFunc {
	return func(server *server) {

		server.router = mux.NewRouter()
		//TODO The time out multiplication time seconds works with const value but not passed in integers. TBD.
		server.srv = &http.Server{
			Handler:      server.router,
			Addr:         host,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		log.Println("Start server on ", server.srv.Addr)
		err := server.srv.ListenAndServe()
		log.Panic(err)
	}
}
var HandlerFunc = func(path string, handler func(w http.ResponseWriter, r *http.Request)) AsyncFunc {
	return func(server *server) {
		log.Println("Setting server: ", server)
		server.router.HandleFunc(path, handler)
	}
}
