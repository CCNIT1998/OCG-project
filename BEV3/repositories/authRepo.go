package repositories

import (
	"errors"
	"ocg-be/database"
	"ocg-be/models"
)

func CheckEmail(email string) error {
	rows, err := database.DB.Query("SELECT email FROM users WHERE email=?", email)
	if err != nil {
		return err
	}
	if rows.Next() {
		return errors.New("exist email")
	}
	return nil
}

func Register(user models.User) error {
	strQuery, err := database.DB.Prepare("INSERT INTO users" +
		"(last_name, first_name, email, password, role_id)" +
		"VALUES (?,?, ?, ?, ?)")
	if err != nil {
		return errors.New("could not create new user")
	}
	_, err = strQuery.Exec(user.Last_name, user.First_name, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}
	return nil
}
func Login(data map[string]string) models.User {
	var user models.User

	row, _ := database.DB.Query("SELECT * FROM users "+
		"WHERE email = ?", data["email"])

	if row.Next() {
		row.Scan(&user.Id, &user.Last_name, &user.First_name, &user.Email, &user.Password, &user.Role)
	}
	return user
}
