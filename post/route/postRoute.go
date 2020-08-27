package route

import (
	postUsecase "github/bhattaraibishal50/blog/post/usecase"

	"github.com/gin-gonic/gin"
)

// PostRoute exports user routes
func PostRoute(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", postUsecase.GetPosts)
	routerGroup.GET("/create", postUsecase.AddPosts)
}
