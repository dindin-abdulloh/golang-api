package main

import (
	"api-library/book"
	"api-library/handlers"
	"fmt"

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
	// db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)

	books, err := bookRepository.FindByID(4)

	fmt.Println(books)

	// for _, book := range books {
	// 	fmt.Println("Tittle :", book.Tittle)
	// }

	r := gin.Default()

	v1 := r.Group("/v1")

	v1.GET("/", handlers.RouteHandler)

	v1.GET("/hello", handlers.HelloHandler)

	v1.GET("/books/:id", handlers.HandlerBooks)

	v1.POST("/books", handlers.BookPostHandler)

	v1.GET("/query", handlers.QueryHandler)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
