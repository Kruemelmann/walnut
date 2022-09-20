package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read token from header
		var token = r.Header.Get("x-access-token")
		json.NewEncoder(w).Encode(r)
		token = strings.TrimSpace(token)

		// check if token belongs to a valid user
		if user, found := validateToken(token); found {
			w.Write([]byte("Token found. Hello " + user + "\n"))
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

//TODO refactor user
//TODO ask login service if token is valid
func validateToken(token string) (user string, status bool) {
	if token == "" {
		return "", false
	}
	if token == "123" {
		return "lukas", true
	}
	return "", false
}
