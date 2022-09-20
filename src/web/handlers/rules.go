package handlers

import (
	"net/http"
)

func RulesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Rules 1, Rules 2, Rules 3"))
	return
}
