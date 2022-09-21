package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	walnut "github.com/kruemelmann/walnut/src/web/genproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read token from header
		var token = r.Header.Get("x-access-token")
		json.NewEncoder(w).Encode(r)
		token = strings.TrimSpace(token)

		// check if token belongs to a valid user
		if user, found := validateToken(token); found {
			w.Write([]byte("Token found. Hello " + user.GetUsername() + "\n"))
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

var authservice_addr = "localhost:50051"

//TODO refactor user
//TODO ask auth service if token is valid
func validateToken(token string) (user *walnut.User, status bool) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(authservice_addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error: did not connect: %v", err)
	}
	defer conn.Close()
	c := walnut.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//Validate Token with the Auth service
	r, err := c.ValidateToken(ctx, &walnut.ValidateTokenRequest{
		Token: token,
	})
	if err != nil {
		log.Printf("error: could not validate token: %v", err)
	}

	return r.GetUser(), r.GetIsValid()
}
