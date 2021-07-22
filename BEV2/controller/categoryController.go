package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ocg-be/models"
	repo "ocg-be/repositories"

	"github.com/gorilla/mux"
)

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categories := repo.GetAllCategories()
	json.NewEncoder(w).Encode(categories)
}

func UpdateCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	category.Id = int64(id)
	err = repo.UpdateCategoryByID(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	category = repo.GetCategoryByID(category.Id)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":         "success",
		"product updated": category,
	})
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = repo.CreateCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "success"})
}
