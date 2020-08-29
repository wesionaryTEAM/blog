package repo

import (
	"context"
	"github/bhattaraibishal50/blog/common"
	"github/bhattaraibishal50/blog/post/model"
	"log"

	"google.golang.org/api/iterator"
)

// PostRepository interface
type PostRepository interface {
	Save(post *model.Post) *model.Post
	FindAll() []*model.Post
}

type firebaseRepo struct{}

var app = common.NewFirebaseApp()
var client = app.GetFirestore()
var ctx = context.Background()

const collectionName string = "posts"

// NewFirebasePostRepository constructor returns the postRepositoory
func NewFirebasePostRepository() PostRepository {
	return &firebaseRepo{}
}

// save function
func (r *firebaseRepo) Save(post *model.Post) *model.Post {
	_, err := client.Collection(collectionName).Doc(post.ID).Set(ctx, map[string]interface{}{
		"ID":          post.ID,
		"Title":       post.Title,
		"Description": post.Description,
		"CreatedAt":   post.CreatedAt,
	})
	if err != nil {
		log.Printf("An error has occurred on adding data: %s", err)
	}
	return post
}

// find all function
func (r *firebaseRepo) FindAll() []*model.Post {
	var posts []*model.Post
	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		var post model.Post
		doc.DataTo(&post) //map to the post struct directory
		posts = append(posts, &post)
	}
	return posts

}

// -----------------------------------------------------------
// mysql repo
// -----------------------------------------------------------

// type mysqlRepo struct{}

// // NewMysqlPostRepository returns struct for mysql
// func NewMysqlPostRepository() PostRepository {
// 	return &mysqlRepo{}
// }

// func (r *mysqlRepo) Save(post *model.Post) *model.Post {
// 	db := common.DbConnection()
// 	fmt.Println(db)
// 	return post
// }
