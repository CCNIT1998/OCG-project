package routes

import (
	"net/http"
	"ocg-be/controller"
	"ocg-be/middlewares"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {

	routeAdmin := r.PathPrefix("/admin").Subrouter()
	routeAdmin.Use(middlewares.IsAuthorized)
	routesPublic(r)
	routesAdmin(routeAdmin)

}
func routesPublic(r *mux.Router) {

	//api-products
	r.HandleFunc("/products", controller.GetProducts).Methods(http.MethodGet) //GET_PRODUCTS
	r.HandleFunc("/products/collections", controller.GetProductCollection).Methods(http.MethodGet)
	r.HandleFunc("/products/{handle}", controller.GetProductByHandle).Methods(http.MethodGet) //GET_PRODUCT_By_Id
	r.HandleFunc("/products", controller.CreateProduct).Methods(http.MethodPost)              //POST_PRODUCT
	r.HandleFunc("/products/{id}", controller.DeleteProductById).Methods(http.MethodDelete)   //DELETE_PRODUCT_By_Id
	r.HandleFunc("/products", controller.UpdateProduct).Methods(http.MethodPatch)             //GET_PRODUCT_By_Id

	//api-image
	r.HandleFunc("/uploads", controller.CreateProductImage).Methods(http.MethodPost) //Upload_images

	image := http.StripPrefix("/images/", http.FileServer(http.Dir("./")))
	r.PathPrefix("/images/").Handler(image) //Prefix-image

	//api-authens
	r.HandleFunc("/login", controller.Login).Methods(http.MethodPost)       //Login
	r.HandleFunc("/register", controller.Register).Methods(http.MethodPost) //Register
	r.HandleFunc("/logout", controller.Logout).Methods(http.MethodPost)     //Logout

	//api-orders

	//collections
	// r.HandleFunc("/collections/:name", controller.GetAllCollections).Methods(http.MethodGet)
	r.HandleFunc("/collections", controller.GetAllCollections).Methods(http.MethodGet)
	r.HandleFunc("/collections", controller.CreateCollection).Methods(http.MethodPut)
	r.HandleFunc("/collections", controller.AddProductToCollection).Methods(http.MethodPost)
	r.HandleFunc("/collects/{id}", controller.RemoveProductFromCollection).Methods(http.MethodDelete)
	r.HandleFunc("/collections/{id}", controller.DeleteCollectionById).Methods(http.MethodDelete)
}
func routesAdmin(r *mux.Router) { //admin/...
	r.HandleFunc("/orders", controller.GetAllOrders).Methods(http.MethodGet)

}
