package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func checkSessionID(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sessionID string

		c, err := r.Cookie(cookieSessionID)
		if err == http.ErrNoCookie {
			// if no cookie is set this will create a new one
			u, _ := uuid.NewRandom()
			sessionID = u.String()
			http.SetCookie(w, &http.Cookie{
				Name:   cookieSessionID,
				Value:  sessionID,
				MaxAge: cookieSessionMaxAge,
			})
		} else if err != nil {
			// if a different error occurs just return
			return
		} else {
			// if a cookie is found just set the value of it as the sessionID
			sessionID = c.Value
		}
		ctx := context.WithValue(r.Context(), ctxSessionID{}, sessionID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
