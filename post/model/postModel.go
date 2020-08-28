package model

// Post struct with the post attributes
type Post struct {
	ID          string
	Title       string ` json:"Title" binding:"required"`
	Description string ` json:"Description" binding:"required"`
	CreatedAt   string
}
