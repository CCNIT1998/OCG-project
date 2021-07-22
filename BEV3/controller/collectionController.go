package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ocg-be/models"
	"ocg-be/repositories"
	"ocg-be/util"
	"strconv"

	"github.com/gorilla/mux"
)

var collectionRepo *repositories.CollectionRepo

func GetAllCollections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 9
	}
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
	products := repositories.Pageniate(&repositories.CollectionRepo{}, page, limit, orderBy, sort, search)
	respon, _ := json.Marshal(products)
	w.WriteHeader(http.StatusOK)
	w.Write(respon)
}

func CreateCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var collection *models.Collection

	err := util.BodyParser(&collection, w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// thêm sản phẩm vào dtb
	collection.Handle = util.ParseNametoHandle(collection.Name)
	collection.Id, err = collectionRepo.Add(collection)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	//chuyển về dạng jsón
	respon, err := json.Marshal(collection)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(respon))
}
func DeleteCollectionById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := collectionRepo.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("delete success"))
}

func AddProductToCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var collectionProduc *models.CollectionProduct

	err := util.BodyParser(&collectionProduc, w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		w.Write([]byte("anh xem lai request gui dung dinh dang khong"))

		return
	}
	err = collectionRepo.AddProductToCollectionDb(collectionProduc.CollectionId, collectionProduc.ProductId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		w.Write([]byte("anh xem lai request gui dung product id voi collection id khong (phai gui nhung cai co ton tai roi)"))

		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("message:ok"))
}

func RemoveProductFromCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := collectionRepo.RemoveProductFromCollection(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("delete success"))
}

// func UpdateProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var product *models.Product
// 	err := util.BodyParser(&product, w, r)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(err.Error()))
// 	}

// 	dbResult, err := productStorage.UpdateProduct(product)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(err.Error()))
// 	}
// 	log.Println(dbResult)

// 	respon, _ := json.Marshal(dbResult)

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(respon))

// }
