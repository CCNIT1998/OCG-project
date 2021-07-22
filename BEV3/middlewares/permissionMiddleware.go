package middlewares

import (
	"fmt"
	"net/http"
	"ocg-be/database"
	"ocg-be/models"
	"ocg-be/util"
	"strconv"
)

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")

		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		Id, err := util.ParseJwt(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		fmt.Println(Id)

		userId, _ := strconv.Atoi(Id)

		var user models.User
		row, err := database.DB.Query("select role_id from users where id = ? ", userId)
		if err != nil {
			fmt.Println(err)
		}
		if row.Next() {
			row.Scan(&user.Role)
		}
		fmt.Println(user.Role)
		if user.Role != 1 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
