package repo

import (
	"context"
	"fmt"
	"github/bhattaraibishal50/blog/common"
	"github/bhattaraibishal50/blog/post/model"
	"log"
	"time"
)

// PostRepository interface
type PostRepository interface {
	Save(post *model.Post) *model.Post
}

// Repo struct implements the interface
type repo struct{}

const collectionName string = "posts"

// NewPostRepository constructor returns the postRepositoory
func NewPostRepository() PostRepository {
	return &repo{}
}

func (r *repo) Save(post *model.Post) *model.Post {
	app := common.InitFirebase()
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	defer client.Close()
	if err != nil {
		log.Printf("An error has occurred on setting client: %s", err)
	}
	results, err := client.Collection(collectionName).Doc("LA").Set(ctx, map[string]interface{}{
		"ID":          post.ID,
		"Title":       post.Title,
		"Description": post.Description,
		"CreatedAt":   time.Now(),
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred on adding data: %s", err)
	}
	fmt.Println(results)
	return post
}
