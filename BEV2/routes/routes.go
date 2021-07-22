package routes

import (
	"net/http"
	"ocg-be/controller"
	"ocg-be/middlewares"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {

	routeAdmin := r.PathPrefix("/admin").Subrouter()
	routeAdmin.Use(middlewares.IsAuthenticated)
	routesPublic(r)
	routesAdmin(routeAdmin)

}
func routesPublic(r *mux.Router) {

	//api-products
	r.HandleFunc("/products", controller.GetProducts).Methods(http.MethodGet)               //GET_PRODUCTS
	r.HandleFunc("/products/{id}", controller.GetProductById).Methods(http.MethodGet)       //GET_PRODUCT_By_Id
	r.HandleFunc("/products", controller.CreateProduct).Methods(http.MethodPost)            //POST_PRODUCT
	r.HandleFunc("/products/{id}", controller.DeleteProductById).Methods(http.MethodDelete) //DELETE_PRODUCT_By_Id
	r.HandleFunc("/products", controller.UpdateProduct).Methods(http.MethodPatch)           //GET_PRODUCT_By_Id

	//api-image
	r.HandleFunc("/uploads", controller.CreateProductImage).Methods(http.MethodPost) //Upload_images

	image := http.StripPrefix("/images/", http.FileServer(http.Dir("./")))
	r.PathPrefix("/images/").Handler(image) //Prefix-image

	//api-authens
	r.HandleFunc("/login", controller.Login).Methods(http.MethodPost)       //Login
	r.HandleFunc("/register", controller.Register).Methods(http.MethodPost) //Register

}
func routesAdmin(r *mux.Router) {
	r.HandleFunc("/logout", controller.Logout).Methods(http.MethodPost)
}
