package config

import (
	"context"
	"errors"
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
		return errors.New("Error firebase new app -> " + err.Error());
	}

	AuthClient, err = app.Auth(ctx);
	if err != nil {
		return errors.New("Error auth firebase client -> " + err.Error());
	}

	return nil;
}