package repo

import (
	"context"
	"github/bhattaraibishal50/blog/common"
	"github/bhattaraibishal50/blog/post/model"
	"log"
)

// PostRepository interface
type PostRepository interface {
	Save(post *model.Post) *model.Post
}

const collectionName string = "posts"

type firebaseRepo struct{}

// NewFirebasePostRepository constructor returns the postRepositoory
func NewFirebasePostRepository() PostRepository {
	return &firebaseRepo{}
}

func (r *firebaseRepo) Save(post *model.Post) *model.Post {
	app := common.InitFirebase()
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	defer client.Close()
	if err != nil {
		log.Printf("An error has occurred on setting client: %s", err)
	}
	_, err = client.Collection(collectionName).Doc(common.GetNewUUID().String()).Set(ctx, map[string]interface{}{
		"ID":          post.ID,
		"Title":       post.Title,
		"Description": post.Description,
		"CreatedAt":   post.CreatedAt,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred on adding data: %s", err)
	}
	return post
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
