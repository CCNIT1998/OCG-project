package models

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Handle      string  `json:"handle"`
	Description string  `json:"description"`
	Price       int64   `json:"price"`
	Trademark   string  `json:"trade_mark"`
	Image       []Image `json:"image"`
}
