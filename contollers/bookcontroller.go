package controllers

import (
	"net/http"

	"github.com/Christomesh/gopostgres/models"
	"github.com/gin-gonic/gin"
)

var BookDB = models.SetUpBookMigration()

func CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := BookDB.Create(&book).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": "book has been added"})

	}
}

func GetBookById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		var book_id = c.Param("book_id")
		if book_id == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "id cannot be empty"})
			return
		}
		err := BookDB.Where("id=?", book_id).First(&book).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "could not get the book"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": book})
	}
}

func GetBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookModels := &[]models.Book{}

		err := BookDB.Find(bookModels).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "could not get books"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "book fetched successfully", "data": bookModels})
	}
}

func DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookModel := models.Book{}
		book_id := c.Param("book_id")
		if book_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "id cannot be empty"})
		}

		err := BookDB.Delete(bookModel, book_id)

		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "could not delete book"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "book delete successfully"})
	}
}
