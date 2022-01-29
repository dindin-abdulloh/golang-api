package book

type BookResponse struct {
	ID          int    `json:"id"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Discount    int    `json:"discount"`
}
