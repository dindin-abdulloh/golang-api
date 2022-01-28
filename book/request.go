package book

import "encoding/json"

type BookRequest struct {
	Tittle      string      `json:"tittle" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	Rating      int         `json:"rating" binding:"required,number"`
	Discount    int         `json:"discount" binding:"required,number"`
}
