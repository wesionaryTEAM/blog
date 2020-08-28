/*
Package usecase is a business logics aread
*/
package usecase

import (
	"github/bhattaraibishal50/blog/common"
	"github/bhattaraibishal50/blog/post/model"
	postRepo "github/bhattaraibishal50/blog/post/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPosts gets the posts
func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "on posts"})
}

// AddPosts add the post
func AddPosts(c *gin.Context) {
	post := model.Post{}
	post.ID = common.GetNewUUID()
	post.Title = "TitleFromCode"
	post.Description = "Description from code"
	postRepo := postRepo.NewFirebasePostRepository()
	res := postRepo.Save(&post)
	c.JSON(http.StatusOK, gin.H{"status": &res})
}
