package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"api-library/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

// GET
func (h *bookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertBookResponse(b)

		booksResponse = append(booksResponse, bookResponse)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

// GET BY ID
func (h *bookHandler) GetBook(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindByID(int(id))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	bookResponse := convertBookResponse(b)
	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

// POST
func (h *bookHandler) CreateBook(ctx *gin.Context) {
	var bookRequest book.BookRequest

	err := ctx.ShouldBindJSON(&bookRequest)

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

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertBookResponse(book),
	})
}

// PUT

func (h *bookHandler) UpdateBook(ctx *gin.Context) {
	var bookRequest book.BookRequest

	err := ctx.ShouldBindJSON(&bookRequest)

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

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertBookResponse(book),
	})
}

// DELETE
func (h *bookHandler) DeleteBook(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindByID(int(id))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	bookResponse := convertBookResponse(b)
	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

// Privat function for convert response
func convertBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Tittle:      b.Tittle,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}
