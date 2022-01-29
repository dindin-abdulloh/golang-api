package main

import (
	"api-library/book"
	"api-library/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)

	// for _, book := range books {
	// 	fmt.Println("Tittle :", book.Tittle)
	// }

	r := gin.Default()

	v1 := r.Group("/v1")

	v1.POST("/books", bookHandler.CreateBook)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// ++++SKEMA++++
// main
// handler
// service
// repository
// db
// mysql
