package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"ocg-be/database"
	"ocg-be/models"
	"reflect"
)

type ProductStorage struct {
}

func (p *ProductStorage) Count(categoryId int, search string) int64 {
	var rows *sql.Rows
	var err error

	if categoryId == 0 {
		rows, err = database.DB.Query("SELECT COUNT(*) FROM products WHERE name LIKE ?", search)
		if err != nil {
			log.Println(err)
		}
	} else {
		rows, err = database.DB.Query("SELECT COUNT(*) FROM products WHERE category_id=? AND name LIKE ?", categoryId, search)
		if err != nil {
			log.Println(err)
		}
	}
	var total int64
	for rows.Next() {
		rows.Scan(&total)
	}

	return total
}
func (p *ProductStorage) TakeByParams(limit int, offset int, categoryId int, orderBy string, sort string, search string) interface{} {
	var products []models.Product
	var rows *sql.Rows
	var err error

	qtext := fmt.Sprintf("SELECT * FROM products WHERE name LIKE ? ORDER BY %s %s LIMIT ? OFFSET ? ", orderBy, sort)
	qtextCate := fmt.Sprintf("SELECT * FROM products WHERE name LIKE ? AND category_id = ? ORDER BY %s %s LIMIT ? OFFSET ? ", orderBy, sort)

	if categoryId == 0 {
		rows, err = database.DB.Query(qtext, search, limit, offset)
	} else {
		rows, err = database.DB.Query(qtextCate, search, categoryId, limit, offset)
	}

	if err != nil {
		log.Println(err)
		return nil
	}

	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Trademark)

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

func (p *ProductStorage) Add(product *models.Product) int {

	_, err := database.DB.Query("INSERT INTO products VALUES (?,?, ?, ?, ?, ?)", product.Id, product.Name, product.Description, product.Price, product.CategoryId, product.Trademark)

	if err != nil {
		product = nil
		fmt.Println(err)
		return 0
	}
	rows, _ := database.DB.Query("SELECT id FROM products ORDER BY id DESC LIMIT 1")

	if rows.Next() {
		err = rows.Scan(&product.Id)
		if err != nil {
			product = nil
			fmt.Println(err)
			return 0
		}
	}

	return product.Id
}

func (p *ProductStorage) GetById(id int) models.Product {
	var product models.Product
	rows, err := database.DB.Query("SELECT * from products WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)

		return product
	}
	if rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Trademark)
		if err != nil {
			fmt.Println(err)
			return product
		}
	}
	var image ImageStorage

	product.Image = image.Take(int64(product.Id))

	return product
}

func (p *ProductStorage) DeleteById(id int) {
	_, err := database.DB.Query("DELETE FROM products WHERE id=?", id)
	if err != nil {
		log.Println(err)
	}
}

func (p *ProductStorage) UpdateProduct(product *models.Product) (*models.Product, error) {
	log.Println(product.Id)
	fmt.Println(reflect.TypeOf(product.Id))
	if product.CategoryId != 0 {
		_, err := database.DB.Query("UPDATE products SET category_id=? WHERE id=?", product.CategoryId, product.Id)
		if err != nil {
			return nil, err
		}
	}
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
