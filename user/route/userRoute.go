package route

import (
	"github/bhattaraibishal50/blog/user/usecase"

	"github.com/gin-gonic/gin"
)

// UserRoute exports user routes
func UserRoute(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", usecase.GetUsers)
}
