package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"ocg-be/database"
	"ocg-be/models"
)

type CollectionRepo struct {
}

func (*CollectionRepo) Count(search string) int64 {
	var rows *sql.Rows
	var err error

	rows, err = database.DB.Query("SELECT COUNT(*) FROM collections WHERE name LIKE ?", search)
	if err != nil {
		log.Println(err)
	}

	var total int64
	for rows.Next() {
		rows.Scan(&total)
	}

	return total
}
func (*CollectionRepo) Take(limit int, offset int, orderBy string, sort string, search string) interface{} {
	var collections []models.Collection
	var rows *sql.Rows
	var err error

	qtext := fmt.Sprintf("SELECT * FROM collections WHERE name LIKE ? ORDER BY %s %s LIMIT ? OFFSET ? ", orderBy, sort)

	rows, err = database.DB.Query(qtext, search, limit, offset)

	if err != nil {
		log.Println(err)
		return nil
	}

	for rows.Next() {
		var collection models.Collection
		err = rows.Scan(&collection.Id, &collection.Name, &collection.Handle, &collection.Thumbnail)

		if err != nil {
			log.Println(err)
			return nil
		}
		collections = append(collections, collection)
	}
	return collections
}
func (*CollectionRepo) Add(collection *models.Collection) (int, error) {

	_, err := database.DB.Query("INSERT INTO collections VALUES (?,?, ?,?)", collection.Id, collection.Name, collection.Handle, collection.Thumbnail)

	if err != nil {
		collection = nil
		fmt.Println(err)
		return 0, errors.New("collection bi trung ten anh oi")
	}
	rows, _ := database.DB.Query("SELECT id FROM collections ORDER BY id DESC LIMIT 1")

	if rows.Next() {
		err = rows.Scan(&collection.Id)
		if err != nil {
			collection = nil
			fmt.Println(err)
			return 0, err
		}
	}

	return int(collection.Id), nil
}

func (*CollectionRepo) Delete(id int) error {
	_, err := database.DB.Query("DELETE FROM collection_product WHERE collection_product.collection_id=?", id)
	if err != nil {
		return err
	}
	_, err = database.DB.Query("DELETE FROM collections WHERE collections.id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func (*CollectionRepo) GetCollectionByHandle(handle string) (models.Collection, error) {
	var collection models.Collection
	rows, err := database.DB.Query("SELECT * from collections WHERE handle = ?", handle)
	if err != nil {
		fmt.Println(err)
		return collection, err
	}
	if rows.Next() {
		err = rows.Scan(&collection.Id, &collection.Name, &collection.Handle, &collection.Thumbnail)

		if err != nil {
			fmt.Println(err)
			return collection, err
		}
	} else {
		return collection, errors.New("khong tim thay collection")
	}
	return collection, nil
}
func (*CollectionRepo) AddProductToCollectionDb(collection_id int, productId int) error {
	qtext, err := database.DB.Prepare(" INSERT INTO collection_product " +
		"(product_id,collection_id) " +
		"VALUES (?,?) ")
	if err != nil {
		return err
	}
	_, err = qtext.Exec(productId, collection_id)
	if err != nil {
		return err
	}
	return nil
}
func (*CollectionRepo) RemoveProductFromCollection(productId int) error {
	_, err := database.DB.Query("DELETE FROM collection_product WHERE collection_product.product_id=?", productId)
	if err != nil {
		return err
	}
	return nil
}
