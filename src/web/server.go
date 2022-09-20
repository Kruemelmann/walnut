package main

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kruemelmann/walnut/src/web/handlers"
)

//Constructor
func NewServer(host, port string) *Server {
	return &Server{
		Host: host,
		Port: port,
	}
}

type Server struct {
	httpserver *http.Server
	Host       string
	Port       string
}

func (s *Server) Start() {
	r := mux.NewRouter().StrictSlash(true)

	secure := r.PathPrefix("/api").Subrouter()
	secure.Use(handlers.JwtVerify)
	secure.HandleFunc("/rules", handlers.RulesHandler)

	r.PathPrefix("/").Handler(getStaticFilesHandler())

	adress := s.Host + ":" + s.Port
	log.Println("Server started on " + adress)
	log.Fatal(http.ListenAndServe(adress, r))
}

//TODO refactor maybe move to own package
func getStaticFilesHandler() http.Handler {
	matches, _ := fs.Glob(FrontendFS, "static/ui/dist")
	if len(matches) != 1 {
		panic("unable to find frontend build files in FrontendFS")
	}
	feRoot, _ := fs.Sub(FrontendFS, matches[0])
	buildHandler := http.FileServer(http.FS(feRoot))
	return buildHandler
}
