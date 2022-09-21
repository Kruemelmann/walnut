package main

import (
	"context"
	"fmt"

	walnut "github.com/kruemelmann/walnut/src/auth/genproto"
)

type authServer struct {
	walnut.UnimplementedAuthServiceServer
}

func (a *authServer) ValidateToken(ctx context.Context, in *walnut.ValidateTokenRequest) (*walnut.ValidateTokenResponse, error) {

	fmt.Printf("Validate Token %s\n", in.GetToken())

	//TODO just for testing
	isValid := false
	if in.GetToken() == "123" {
		isValid = true
	}

	return &walnut.ValidateTokenResponse{
		//TODO validate with the db
		IsValid: isValid,
		User: &walnut.User{
			Username: "lukas",
			Password: "123",
			Role:     walnut.Role_ADMIN,
		},
	}, nil
}

func (a *authServer) Login(ctx context.Context, in *walnut.LoginRequest) (*walnut.LoginResponse, error) {
	//TODO implement
	return &walnut.LoginResponse{
		Successful: true,
		Token:      "123",
	}, nil
}
