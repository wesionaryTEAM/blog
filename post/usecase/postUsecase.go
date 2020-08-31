/*
Package usecase is a business logics aread
*/
package usecase

import (
	"github/bhattaraibishal50/blog/common"
	"github/bhattaraibishal50/blog/post/model"
	postRepo "github/bhattaraibishal50/blog/post/repo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetPosts gets the posts
func GetPosts(c *gin.Context) {
	postRepo := postRepo.NewFirebasePostRepository()
	posts := postRepo.FindAll()
	c.JSON(http.StatusOK, gin.H{"data": &posts})
}

// AddPosts add the post
func AddPosts(c *gin.Context) {
	// for json type data we need to bind with the struct references https://mholt.github.io/json-to-go/
	//JSON BIND
	post := model.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.ID = common.GetNewUUID().String()
	post.CreatedAt = time.Now().String()
	c.Bind(post)
	postRepo := postRepo.NewFirebasePostRepository()
	res := postRepo.Save(&post)
	c.JSON(http.StatusOK, gin.H{"data": &res})
}
