package repositories

import (
	"ocg-be/database"
	"ocg-be/models"
)

func GetAllOrders() []models.Order {

	rows, _ := database.DB.Query("SELECT id,create_at,address,phone_number,total_price,order_notes " +
		"FROM orders")
	var orders []models.Order
	for rows.Next() {
		var order models.Order
		rows.Scan(&order.Id, &order.CreatedAt, &order.Address, &order.PhoneNumber, &order.TotalPrice, &order.OrderNote)
		order.OrderItems = getOrderItem(order.Id)
		orders = append(orders, order)
	}
	return orders
}

func getOrderItem(id int64) []models.OderItem {
	rows, _ := database.DB.Query("SELECT product_id,quantity FROM order_items "+
		"WHERE order_id = ?", id)
	var orderItems []models.OderItem
	for rows.Next() {
		var orderItem models.OderItem
		rows.Scan(&orderItem.ProductId, &orderItem.Quantity)
		orderItems = append(orderItems, orderItem)
	}
	return orderItems
}

func GetOrderByID(id int64) models.Order {

	rows, _ := database.DB.Query("SELECT id,create_at,address,phone_number,total_price,order_notes,status FROM orders "+
		"WHERE id = ?", id)
	var order models.Order
	if rows.Next() {
		rows.Scan(&order.Id, &order.CreatedAt, &order.Address, &order.PhoneNumber, &order.TotalPrice, &order.OrderNote, &order.Status)
	}
	order.OrderItems = getOrderItem(order.Id)
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

