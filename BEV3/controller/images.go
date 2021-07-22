package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"ocg-be/models"
	"ocg-be/repositories"
	"strconv"
)

var imageStorage repositories.ImageStorage

func CreateProductImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	product_id, _ := strconv.Atoi(r.FormValue("product_id"))

	if product_id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing product id"))
		return
	}

	files := r.MultipartForm.File["image"]
	totalFile, _ := repositories.UploadFile(files)

	var image models.Image
	for _, filename := range totalFile {
		image.Url = filename
		err = imageStorage.AddImage(int64(product_id), image.Url)
		if err != nil {
			log.Println(err)
			return
		}
	}
	images := imageStorage.GetImageByProductId(int64(product_id))
	respon, _ := json.Marshal(images)
	w.WriteHeader(http.StatusOK)
	w.Write(respon)
}
