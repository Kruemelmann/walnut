package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	walnut "github.com/kruemelmann/walnut/src/web/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	t, err := login(u.Username, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	} else {
		u.Token = t
	}
	u.Password = ""

	res, err := json.Marshal(u)
	if err != nil {
		//TODO dont panic and handle error
		panic(err)
	}

	w.Write(res)
	return
}

//TODO refactor to own layer
func login(username, password string) (string, error) {
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
	r, err := c.Login(ctx, &walnut.LoginRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Printf("error: could not login: %v", err)
	}

	if r.GetSuccessful() {
		return r.GetToken(), nil
	} else {
		return "", errors.New("unable to login user " + err.Error())
	}
}
