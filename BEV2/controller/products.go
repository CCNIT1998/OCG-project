package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"ocg-be/models"
	"ocg-be/repositories"
	"ocg-be/util"
	"strconv"

	"github.com/gorilla/mux"
)

var productStorage *repositories.ProductStorage

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 9
	}
	categoryId, _ := strconv.Atoi(r.URL.Query().Get("category"))
	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "asc"
	}
	orderBy := r.URL.Query().Get("order")
	if orderBy == "" {
		orderBy = "id"
	}
	search := r.FormValue("search")
	search = "%" + search + "%"
	products := repositories.Pageniate(&repositories.ProductStorage{}, page, limit, categoryId, orderBy, sort, search)
	respon, _ := json.Marshal(products)
	w.WriteHeader(http.StatusOK)
	w.Write(respon)

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product *models.Product

	err := util.BodyParser(&product, w, r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	// thêm sản phẩm vào dtb
	product.Id = productStorage.Add(product)
	//chuyển về dạng jsón
	respon, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	//trả kết quả
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respon))
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	dbResult := productStorage.GetById(id)
	respon, _ := json.Marshal(dbResult)
	//trả kết quả
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respon))
}

func DeleteProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	imageStorage.DeleteImageByProductId(id)
	productStorage.DeleteById(id)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("delete success"))

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product *models.Product
	err := util.BodyParser(&product, w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	dbResult, err := productStorage.UpdateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	log.Println(dbResult)

	respon, _ := json.Marshal(dbResult)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respon))

}
