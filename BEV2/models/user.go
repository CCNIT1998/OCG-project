package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id         int64  `json:"id"`
	Last_name  string `json:"last_name"`
	First_name string `json:"first_name"`
	Email      string `json:"email"`
	Password   []byte `json:"-"`
	Role       int64  `json:"role"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
