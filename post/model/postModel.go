package model

import "github.com/google/uuid"

// Post struct with the post attributes
type Post struct {
	ID          uuid.UUID
	Title       string
	Description string
	CreatedAt   string
}
