package routes

import (
	controller "github.com/Christomesh/gopostgres/contollers"
	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine) {
	book := router.Group("/api/v1/books")

	{
		book.GET("/", controller.GetBooks())
		book.POST("/", controller.CreateBook())
		book.GET("/:book_id", controller.GetBookById())
		book.DELETE("/:book_id", controller.DeleteBook())
	}
}
