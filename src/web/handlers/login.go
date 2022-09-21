package handlers

import (
	"encoding/json"
	"net/http"
)

type user struct {
	Username string
	Password string
	Token    string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u user

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//TODO check user with auth service and return with token to gui

	res, err := json.Marshal(u)
	if err != nil {
		//TODO dont panic and handle error
		panic(err)
	}

	w.Write(res)
	return
}
