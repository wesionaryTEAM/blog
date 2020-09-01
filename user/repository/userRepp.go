package repository

import (
	"context"
	"fmt"
	"github/bhattaraibishal50/blog/common"
)

// UserRepository interface
type UserRepository interface {
	Login()
}

type firebaseAuthRepo struct{}

var fb = common.NewFirebaseApp()
var fbAuth = fb.GetFirebaseAuth()
var ctx = context.Background()

// NewFirebaseAuthRepo constrictore returns user repos
func NewFirebaseAuthRepo() UserRepository {
	return &firebaseAuthRepo{}
}

// Login func
// The “Firebase Admin SDK” is for administrative actions only (fetching user data, changing the user’s email, etc) without their existing password. 
// If you want to sign in as a user, you need to use the Firebase Authentication SDK.
func (*firebaseAuthRepo) Login() {
	// checks the user by email and the password
	usersRecord, err := fbAuth.GetUser(ctx, "5ZamU26mXsOYmCrankcEGTujgsj2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(usersRecord)
}


// GO tokenization.
