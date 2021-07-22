package models

type Order struct {
	Id          int64      `json:"id"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	OrderItems  []OderItem `json:"order_items"`
	TotalPrice  int64      `json:"total_price"`
	CreatedAt   string     `json:"created_at"`
	// User        []User     `json:"user"`
	OrderNote string `json:"order_note"`
	Status    bool   `json:"status"`
}
