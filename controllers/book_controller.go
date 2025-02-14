package controllers

import (
	"net/http"

	"gin_gorm/repository"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	books, err := repository.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}