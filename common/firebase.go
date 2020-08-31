// Package common  - common packs
// References : https://github.com/firebase/firebase-admin-go
package common

import (
	"log"
	"path/filepath"

	"firebase.google.com/go/auth"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// Firebase interface
type Firebase interface {
	GetFirestore() *firestore.Client
	GetFirebaseAuth() *auth.Client
}

type fb struct{}

var ctx = context.Background()

// NewFirebaseApp  returns Firebase struct implemnting firebase interface
func NewFirebaseApp() Firebase {
	return &fb{}
}

// GetFirestore -  returns firestore.Client
func (*fb) GetFirestore() *firestore.Client {
	app := InitFirebase()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("An error has occurred on setting client: %s", err)
	}
	return client
}

// GetFirebaseAuth -  returns auth.Client
func (*fb) GetFirebaseAuth() *auth.Client {
	app := InitFirebase()
	auth, err := app.Auth(ctx)
	if err != nil {
		log.Printf("An error has occurred on setting client: %s", err)
	}
	return auth
}

// InitFirebase - initlise the firebase & returns firebase.app
func InitFirebase() *firebase.App {
	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json") //shows the full path
	if err != nil {
		panic("Unable to load serviceAccountKey.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	config := &firebase.Config{ProjectID: "blog-goland"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		panic(err.Error())
	}
	return app
}
