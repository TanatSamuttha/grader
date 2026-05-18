package config

import (
	"context"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var AuthClient *auth.Client;

func InitFirebase() error {
	cred := os.Getenv("FIREBASE_CREDENTIALS");

	opt := option.WithCredentialsJSON([]byte(cred));
	ctx := context.Background();
	app, err := firebase.NewApp(ctx, nil, opt);
	if err != nil {
		return err;
	}

	AuthClient, err = app.Auth(ctx);
	if err != nil {
		return err;
	}

	return nil;
}