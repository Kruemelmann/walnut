package main

import (
	"log"
	"net/http"
)

func (ws *webServer) testHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("test handler called")
}
