package common

import (
	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// InitFirebase - Opening a database and save the reference to `Database` struct.
func InitFirebase() *firebase.App {
	opt := option.WithCredentialsFile("/home/bishal/go/src/github/bhattaraibishal50/blog/serviceAccountKey.json")
	config := &firebase.Config{ProjectID: "blog-goland"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		panic(err.Error())
	}
	return app
}
