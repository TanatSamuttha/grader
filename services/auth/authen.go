package main

import (
	"auth/config"
	"context"
	"fmt"
)

func Authen(token string) (string, error) {
	decodedToken, err := config.AuthClient.VerifyIDToken(
		context.Background(),
		token,
	)

	if err != nil {
		return "", err;
	}

	fmt.Println(decodedToken);

	return "token", err;
}