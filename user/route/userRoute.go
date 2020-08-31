package route

import (
	userUsecase "github/bhattaraibishal50/blog/user/usecase"

	"github.com/gin-gonic/gin"
)

// UserRoute exports user routes
func UserRoute(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", userUsecase.GetUsers)
	routerGroup.POST("/signup", userUsecase.Signup)
	routerGroup.POST("/login", userUsecase.Login)
}
