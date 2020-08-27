package middleware

import (
	postRouter "github/bhattaraibishal50/blog/post/route"
	userRouter "github/bhattaraibishal50/blog/user/route"

	"github.com/gin-gonic/gin"
)

// SetRoutes sets routes
func SetRoutes(router *gin.Engine) {
	// user router
	usersGroup := router.Group("/users")
	userRouter.UserRoute(usersGroup)
	// post router
	postGroup := router.Group("/posts")
	postRouter.PostRoute(postGroup)
}
