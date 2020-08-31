/*
Package usecase is a business logics aread
*/
package usecase

import (
	"github/bhattaraibishal50/blog/common"
	"github/bhattaraibishal50/blog/user/model"
	userRepo "github/bhattaraibishal50/blog/user/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GetUsers gets the users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "on users"})
}

// Login godoc
// @Summary Login user
// @Description Login user by email and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param Body body userModel.LoginUser true "Login user"
// @Failure 400 {object} common.HTTPError
// @Failure 404 {object} common.HTTPError
// @Failure 500 {object} common.HTTPError
// @Router /users/login [post]
func Login(c *gin.Context) {
	var loginUser model.LoginUser
	err := c.ShouldBindBodyWith(&loginUser, binding.JSON)
	if err != nil {
		common.NewHTTPError(c, http.StatusBadRequest, err)
		return
	}
	userRepo := userRepo.NewFirebaseAuthRepo()
	userRepo.Login()
	c.JSON(http.StatusOK, gin.H{"data": &loginUser})
}

// Signup signup the user
func Signup(c *gin.Context) {

}
