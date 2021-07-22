package models

type Order struct {
	Id          int64       `json:"id"`
	Last_name   string      `json:"last_name"`
	First_name  string      `json:"first_name"`
	Email       string      `json:"email"`
	PhoneNumber string      `json:"phone_number"`
	Address     string      `json:"address"`
	OrderItems  []OrderItem `json:"order_items"`
	TotalPrice  int64       `json:"total_price"`
	CreatedAt   string      `json:"created_at"`
	OrderNote   string      `json:"order_note"`
	Status      bool        `json:"status"`
}
