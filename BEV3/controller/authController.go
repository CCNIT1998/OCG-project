package controller

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"
	"strconv"
	"time"

	"ocg-be/database"
	"ocg-be/models"
	"ocg-be/repositories"
	"ocg-be/util"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if data["password"] != data["password_confirm"] {
		http.Error(w, "password match incorrect", http.StatusBadRequest)
		return
	}
	user := models.User{
		Last_name:  data["last_name"],
		First_name: data["first_name"],
		Email:      data["email"],
		Role:       2,
	}

	user.SetPassword(data["password"])         //ma hoa
	err := repositories.CheckEmail(user.Email) //check email
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	repositories.Register(user) //Luu user
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Register successfully",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(data)

	user := repositories.Login(data)
	if user.Id == 0 {
		http.Error(w, "email incorrect", http.StatusBadRequest)
		return
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		http.Error(w, "password incorrect", http.StatusBadRequest)
		return
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	})
	response, _ := json.Marshal(user)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))

}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "logout successfully",
	})
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := util.ParseJwt(cookie.Value)
	var user models.User
	row, _ := database.DB.Query("select u.id, u.last_name, u.email, r.name "+
		"from users u join roles r on r.id = u.role_id where u.id = ? ", id)
	if row.Next() {
		row.Scan(&user.Id, &user.Last_name, &user.Email, &user.Role)
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateInfo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := util.ParseJwt(cookie.Value)

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("update users set last_name = ?, email = ? "+
		"where id = ?", user.Last_name, user.Email, id)

	if err != nil {
		http.Error(w, "can not update information", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "update information successfully",
	})
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if data["password"] != data["password_confirm"] {
		http.Error(w, "password confirm incorrect", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := util.ParseJwt(cookie.Value)
	userId, _ := strconv.Atoi(id)
	user := models.User{
		Id: int64(userId),
	}
	user.SetPassword(data["password"])

	_, err = database.DB.Exec("update users set password = ? "+
		"where id = ?", user.Password, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "update password successfully",
	})
}
