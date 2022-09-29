package main

import (
	"os"

	"github.com/Christomesh/gopostgres/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.BookRoute(router)

	router.Run(":" + port)

}
