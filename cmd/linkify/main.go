package main

import (
	"fmt"
	"log"

	"linkify/internal/handlers/router"

	docs "linkify/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title		Linkify API
//	@version	1.0
//	@host		localhost:8080

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	fmt.Println("Starting...")
	r := router.New()

	docs.SwaggerInfo.BasePath = "/api/v1/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}
