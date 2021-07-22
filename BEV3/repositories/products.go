package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"ocg-be/database"
	"ocg-be/models"
	"reflect"
)

type ProductStorage struct {
}

type RequestGetProductByCollectionId struct {
	CollectionId int    `json:"collection_id"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
	Page         int    `json:"page"`
	Sort         string `json:"sort"`
	Order        string `json:"order"`
	Search       string `json:"search"`
}

func (p *ProductStorage) Count(search string) int64 {
	var rows *sql.Rows
	var err error

	rows, err = database.DB.Query("SELECT COUNT(*) FROM products WHERE name LIKE ?", search)
	if err != nil {
		log.Println(err)
	}

	var total int64
	for rows.Next() {
		rows.Scan(&total)
	}

	return total
}
func (p *ProductStorage) Take(limit int, offset int, orderBy string, sort string, search string) interface{} {
	var products []models.Product
	var rows *sql.Rows
	var err error

	qtext := fmt.Sprintf("SELECT * FROM products WHERE name LIKE ? ORDER BY %s %s LIMIT ? OFFSET ? ", orderBy, sort)

	rows, err = database.DB.Query(qtext, search, limit, offset)

	if err != nil {
		log.Println(err)
		return nil
	}

	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Handle, &product.Description, &product.Price, &product.Trademark)

		if err != nil {
			log.Println(err)
			return nil
		}

		var image ImageStorage
		product.Image = image.Take(int64(product.Id))
		products = append(products, product)
	}
	return products

}

func (p *ProductStorage) Add(product *models.Product) (int, error) {

	_, err := database.DB.Query("INSERT INTO products VALUES (?,?, ?, ?, ?, ?)", product.Id, product.Name, product.Handle, product.Description, product.Price, product.Trademark)

	if err != nil {
		product = nil
		fmt.Println(err)

		return 0, errors.New("hinh nhu bi trung ten san pham anh oi")
	}
	rows, _ := database.DB.Query("SELECT id FROM products ORDER BY id DESC LIMIT 1")

	if rows.Next() {
		err = rows.Scan(&product.Id)
		if err != nil {
			product = nil
			fmt.Println(err)
			return 0, err
		}
	}

	return product.Id, nil
}

func (p *ProductStorage) GetByHandle(handle string) (models.Product, error) {
	var product models.Product
	rows, err := database.DB.Query("SELECT * from products WHERE handle = ?", handle)
	if err != nil {
		fmt.Println(err)
		return product, err
	}
	if rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Handle, &product.Description, &product.Price, &product.Trademark)
		if err != nil {
			fmt.Println(err)
			return product, err
		}
	} else {
		return product, errors.New("khong tim thay san pham")
	}
	var image ImageStorage

	product.Image = image.Take(int64(product.Id))

	return product, nil
}

func (p *ProductStorage) DeleteById(id int) error {
	_, err := database.DB.Query("DELETE FROM collection_product WHERE product_id=?", id)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = database.DB.Query("DELETE FROM products WHERE id=?", id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func (p *ProductStorage) UpdateProduct(product *models.Product) (*models.Product, error) {
	log.Println(product.Id)
	fmt.Println(reflect.TypeOf(product.Id))

	if product.Price != 0 {
		log.Println("price")
		_, err := database.DB.Query("UPDATE products SET price=? WHERE id=?", product.Price, product.Id)
		if err != nil {
			return nil, err
		}

	}
	if product.Description != "" {
		_, err := database.DB.Query("UPDATE products SET description=? WHERE id=?", product.Description, product.Id)
		if err != nil {
			return nil, err
		}
	}
	if product.Name != "" {
		log.Println("name")
		_, err := database.DB.Query("UPDATE products SET name=? WHERE id=?", product.Name, product.Id)
		if err != nil {
			return nil, err
		}
	}
	if product.Trademark != "" {
		_, err := database.DB.Query("UPDATE products SET trade_mark=? WHERE id=?", product.Trademark, product.Id)
		if err != nil {
			return nil, err
		}
	}

	return product, nil

}
func (p *RequestGetProductByCollectionId) Count(search string) int64 {
	var rows *sql.Rows
	var err error

	rows, err = database.DB.Query("SELECT COUNT(*) FROM `products` LEFT JOIN collection_product ON collection_product.product_id=products.id WHERE collection_product.collection_id= ?", p.CollectionId)
	if err != nil {
		log.Println(err)
	}

	var total int64
	for rows.Next() {
		rows.Scan(&total)
	}

	return total
}

func (p *RequestGetProductByCollectionId) Take(limit int, offset int, orderBy string, sort string, search string) interface{} {
	var products []models.Product
	var rows *sql.Rows
	var err error
	qtext := fmt.Sprintf("SELECT products.* FROM products LEFT JOIN collection_product ON collection_product.product_id=products.id WHERE collection_product.collection_id= ? ORDER BY %s %s LIMIT ? OFFSET ? ", orderBy, sort)

	rows, err = database.DB.Query(qtext, p.CollectionId, limit, offset)

	if err != nil {
		log.Println(err)
		return nil
	}

	for rows.Next() {

		var product models.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Handle, &product.Description, &product.Price, &product.Trademark)

		if err != nil {
			log.Println(err)
			return nil
		}

		var image ImageStorage
		product.Image = image.Take(int64(product.Id))
		products = append(products, product)
	}
	return products
}
