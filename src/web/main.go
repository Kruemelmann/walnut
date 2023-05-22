package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	host = "localhost"
	port = "8000"

	cookieSessionID     = "walnut_session-id"
	cookieSessionMaxAge = 60 * 60 * 10 // 10 hours
)

type ctxSessionID struct{}

//TODO add the grpc bindings
type webServer struct{}

func main() {
	//ctx := context.Background()

	ws := new(webServer)

	r := mux.NewRouter()
	r.HandleFunc("/", ws.testHandler).Methods(http.MethodGet, http.MethodHead)

	var handler http.Handler = r
	handler = checkSessionID(handler)

	adress := host + ":" + port
	log.Println("Server started on " + adress)
	log.Fatal(http.ListenAndServe(adress, handler))
}
