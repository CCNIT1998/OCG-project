package repositories

import (
	"ocg-be/database"
	"ocg-be/models"
)

func GetAllCategories() []models.Category {
	rows, _ := database.DB.Query("SELECT * " +
		"FROM categories")
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		rows.Scan(&category.Id, &category.Name)
		categories = append(categories, category)
	}
	return categories
}

func GetCategoryByID(id int64) models.Category {

	rows, _ := database.DB.Query("SELECT * FROM categories "+
		"WHERE id = ?", id)
	var category models.Category
	if rows.Next() {
		rows.Scan(&category.Id, &category.Name)
	}
	return category
}

func UpdateCategoryByID(category models.Category) error {

	strQuery, err := database.DB.Prepare("UPDATE categories " +
		"SET name = ? WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	strQuery.Exec(category.Name, category.Id)
	return nil
}

func CreateCategory(category models.Category) error {

	row, err := database.DB.Query("SELECT MAX(id) FROM categories")

	if err != nil {
		panic(err.Error())
	}
	if row.Next() {
		row.Scan(&category.Id)
	} else {
		category.Id = 0
	}
	category.Id += 1
	strQuery, err := database.DB.Prepare("INSERT INTO categories " +
		"(name)" + "VALUES (?)")
	if err != nil {
		panic(err.Error())
	}

	strQuery.Exec(category.Name)
	return nil
}

// func DeleteProduct(id int64)error{

// 	strQuery, err := database.DB.Prepare("DELETE FROM products " +
// 		"where id = ?")
// 	if err != nil{
// 		panic(err.Error())
// 	}
// 	strQuery.Exec(id)
// 	return nil
// }
