package repositories

import (
	"ocg-be/database"
	"ocg-be/models"
	// "fmt"
)

func CreateNewUser(user models.User) error {

	row, err := database.DB.Query("SELECT MAX(id) FROM users")

	if err != nil {
		panic(err.Error())
	}
	if row.Next() {
		row.Scan(&user.Id)
	} else {
		user.Id = 0
	}

	user.Id += 1

	strQuery, err := database.DB.Prepare("INSERT INTO users " +
		"(id,last_name, first_name, email, password, role_id) " +
		"VALUES (?,?, ?, ?, ?, 2)")

	if err != nil {
		panic(err.Error())
	}
	// user.SetPassword(string(user.Password))
	strQuery.Exec(user.Last_name, user.First_name, user.Email, user.Password, 2)
	return nil

}

func GetAllUsers() []models.User {
	rows, _ := database.DB.Query("SELECT * " + "FROM users")
	var users []models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.Last_name, &user.First_name, &user.Email, &user.Password, &user.Role)
		users = append(users, user)
	}
	return users
}

func GetUserById(id int64) models.User {
	var user models.User
	row, _ := database.DB.Query("select * from users where id = ? ", id)
	if row.Next() {
		row.Scan(&user.Id, &user.Last_name, &user.First_name, &user.Email, &user.Password, &user.Role)
	}
	return user
}

func UpdateUserById(user models.User) error {
	// var roleId int64
	// if user.Role == "ADMIN" {
	// 	roleId = 1
	// } else {
	// 	roleId = 2
	// }
	strQuery, err := database.DB.Prepare("UPDATE users " +
		"SET last_name = ?, first_name = ?, email = ?, password = ?" + "WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	strQuery.Exec(user.Last_name, user.First_name, user.Email, user.Password, user.Id)
	return nil
}

func DeleteUserById(id int64) error {

	strQuery, err := database.DB.Prepare("DELETE FROM users " + "where id = ?")
	if err != nil {
		panic(err.Error())
	}
	strQuery.Exec(id)
	return nil

}
