package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ocg-be/models"
	repo "ocg-be/repositories"

	"github.com/gorilla/mux"
)

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders := repo.GetAllOrders()
	json.NewEncoder(w).Encode(orders)
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	order := repo.GetOrderByID(int64(id))
	json.NewEncoder(w).Encode(order)
}

func UpdateStatusOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	order.Id = int64(id)
	err = repo.UpdateStatusOrder(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order = repo.GetOrderByID(order.Id)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":       "success",
		"order updated": order.Status,
	})
}
