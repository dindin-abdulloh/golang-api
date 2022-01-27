package book

import "encoding/json"

type BooksInput struct {
	Title string      `json: "tittle" binding: "required"`
	Price json.Number `json: "price" binding: "required,number"`
}
