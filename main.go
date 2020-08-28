package main

import (
	"fmt"
	"github/bhattaraibishal50/blog/common"
	"github/bhattaraibishal50/blog/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/thinkerou/favicon"
)

func main() {
	//setting up the application
	fmt.Println("UUID ::", common.GetNewUUID())
	r := gin.Default()
	r.Use(favicon.New("./public/favicon.ico")) // set favicon middleware

	// set routes
	middleware.SetRoutes(r)
	PORT := ":8080"
	r.Run(PORT) //listen and serve on 0.0.0.0:8080 by default
}
