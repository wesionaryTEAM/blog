package main

import (
	"fmt"
	"github/bhattaraibishal50/blog/common"
	"github/bhattaraibishal50/blog/docs"
	"github/bhattaraibishal50/blog/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/thinkerou/favicon"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


// @contact.name readytowork
// @contact.url http://readytowork.jp
// @contact.email redaytowork@info.jp

func main() {
	//setting up the application
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Blog Application"
	docs.SwaggerInfo.Description = "This is a sample server for Blog Application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "https://8080-c6fca9c0-4e36-41d2-86db-511a4134a8bf.ws-us02.gitpod.io"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	fmt.Println("UUID ::", common.GetNewUUID())
	r := gin.Default()
	r.Use(favicon.New("./public/favicon.ico")) // set favicon middleware

	// set routes
	middleware.SetRoutes(r)
	PORT := ":8080"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(PORT) //listen and serve on 0.0.0.0:8080 by default
}
