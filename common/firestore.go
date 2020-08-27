package common

import (
	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// InitFirebase - Opening a database and save the reference to `Database` struct.
func InitFirebase() *firebase.App {
	opt := option.WithCredentialsFile("../serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err.Error())
	}
	return app
}
