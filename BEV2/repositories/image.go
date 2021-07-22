package repositories

import (
	"log"
	"ocg-be/database"
	"ocg-be/models"
)

type ImageStorage struct {
}

func (*ImageStorage) Take(ProductId int64) []models.Image {
	var images []models.Image
	rows, err := database.DB.Query("SELECT url FROM images WHERE product_id = ?", ProductId)
	if err != nil {
		log.Println(err)
		return nil
	}
	var image models.Image
	for rows.Next() {
		rows.Scan(&image.Url)
		images = append(images, image)
	}
	return images
}

func (*ImageStorage) AddImage(productId int64, url string) error {
	_, err := database.DB.Query("INSERT INTO images VALUES (?,?)", productId, url)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}
func (*ImageStorage) GetImageByProductId(productId int64) []models.Image {
	var images []models.Image
	rows, err := database.DB.Query("SELECT url FROM `images` WHERE product_id=?", productId)
	if err != nil {
		log.Println(err)
		return nil
	}
	var image models.Image
	for rows.Next() {
		err = rows.Scan(&image.Url)
		if err != nil {
			log.Println(err)
			return nil
		}
		images = append(images, image)
	}
	return images
}
func (*ImageStorage) DeleteImageByProductId(productId int) {
	_, err := database.DB.Query("DELETE FROM `images` WHERE product_id=?", productId)
	if err != nil {
		log.Println(err)
		return
	}
}
