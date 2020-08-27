/*
Package usecase is a business logics aread
*/
package usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers gets the users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "on users"})
}
