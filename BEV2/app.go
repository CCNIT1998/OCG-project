package main

import (
	"fmt"
	"log"
	"net/http"
	"ocg-be/database"
	"ocg-be/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	r := mux.NewRouter()

	routes.Setup(r)

	handleCross := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Origin", "Accept"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH"}),
		handlers.AllowedOrigins([]string{"http://localhost:8080"}),
		handlers.AllowCredentials(),
	)

	host := fmt.Sprintf(":%d", 3000) //config
	fmt.Println("http://localhost" + host)
	log.Fatal(http.ListenAndServe(host, handleCross(r)))

}
