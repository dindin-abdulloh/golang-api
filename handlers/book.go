package handlers

import (
	"fmt"
	"net/http"

	"api-library/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RouteHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func HelloHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "INI HELLO",
	})
}

func HandlerBooks(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")

	ctx.JSON(http.StatusOK, gin.H{"title": title})
}

// POST
func BookPostHandler(ctx *gin.Context) {
	var bookInput book.BooksInput

	err := ctx.ShouldBindJSON(&bookInput)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

	}

	ctx.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
