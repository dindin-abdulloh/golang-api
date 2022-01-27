package book

import (
	"time"
)

type Book struct {
	ID          int
	Tittle      string
	Description string
	Price       int
	Rating      int
	Discount    int
	CreatedAt   time.Time
	UpdateAt    time.Time
}
