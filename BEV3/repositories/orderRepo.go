package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"ocg-be/database"
	"ocg-be/models"
)

type OrderStorage struct {
}

func (o *OrderStorage) Count(search string) int64 {
	var rows *sql.Rows
	var err error

	rows, err = database.DB.Query("SELECT COUNT(*) FROM orders WHERE address LIKE ?", search)
	if err != nil {
		log.Println(err)
	}

	var total int64
	for rows.Next() {
		rows.Scan(&total)
	}

	return total
}
func (o *OrderStorage) Take(limit int, offset int, orderBy string, sort string, search string) interface{} {
	qtext := "SELECT " +
		"orders.Id,orders.create_at,orders.address,orders.phone_number,orders.total_price,orders.order_notes,orders.status,users.first_name,users.last_name,users.email " +
		"FROM orders  JOIN users ON orders.user_id = users.id"

	rows, err := database.DB.Query(qtext)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	var orders []models.Order

	for rows.Next() {
		var order models.Order
		err = rows.Scan(&order.Id, &order.CreatedAt,
			&order.Address, &order.PhoneNumber,
			&order.TotalPrice, &order.OrderNote,
			&order.Status, &order.First_name,
			&order.Last_name, &order.Email)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println(order.Id)
		order.OrderItems, err = getOrderItem(order.Id)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		orders = append(orders, order)
	}
	return orders
}

func getOrderItem(id int64) ([]models.OrderItem, error) {
	rows, err := database.DB.Query("SELECT product_id,quantity FROM order_items "+
		"WHERE order_id = ?", id)
	if err != nil {
		return nil, err
	}
	var orderItems []models.OrderItem
	for rows.Next() {
		var orderItem models.OrderItem
		rows.Scan(&orderItem.ProductId, &orderItem.Quantity)
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}

func GetOrderByID(id int64) models.Order {
	var err error
	rows, _ := database.DB.Query("SELECT id,create_at,address,phone_number,total_price,order_notes,status FROM orders "+
		"WHERE id = ?", id)
	var order models.Order
	if rows.Next() {
		rows.Scan(&order.Id, &order.CreatedAt, &order.Address, &order.PhoneNumber, &order.TotalPrice, &order.OrderNote, &order.Status)
	}
	order.OrderItems, err = getOrderItem(order.Id)
	if err != nil {
		return models.Order{}
	}
	return order
}

func UpdateStatusOrder(order models.Order) error {

	_, err := database.DB.Exec("UPDATE orders "+
		"SET status = ? "+
		"WHERE id = ?", order.Status, order.Id)

	if err != nil {
		panic(err.Error())
		// return err
	}

	return nil
}
