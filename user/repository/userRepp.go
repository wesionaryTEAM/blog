package repository

import (
	"context"
	"github/bhattaraibishal50/blog/common"
	"log"
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
func (*firebaseAuthRepo) Login() {
	usersRecord, err := fbAuth.GetUserByEmail(ctx, "bishal.bhattarai@readytowork.jp")
	log.Printf("Successfully fetched user data: %v\n", *usersRecord, err)
}
