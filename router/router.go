package router

import (
	"coutarel/goshop/database"
	"coutarel/goshop/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, userService *database.UserService, productService *database.ProductService, authService *database.AuthService) {

	userHandler := handlers.NewUserHandler(*userService)
	productHandler := handlers.NewProductHandler(*productService)
	authHandler := handlers.NewAuthHandler(*authService)

	apiRoutes := r.PathPrefix("/api").Subrouter()

	apiRoutes.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	apiRoutes.HandleFunc("/auth/signin", authHandler.Signin).Methods("POST")

	apiRoutes.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")
	apiRoutes.HandleFunc("/users/{id}", userHandler.Update).Methods("PATCH")
	apiRoutes.HandleFunc("/users/{id}", userHandler.GetById).Methods("GET")
	apiRoutes.HandleFunc("/users", userHandler.GetAll).Methods("GET")
	apiRoutes.HandleFunc("/users", userHandler.Create).Methods("POST")

	apiRoutes.HandleFunc("/products/{id}", productHandler.GetById).Methods("GET")
	apiRoutes.HandleFunc("/products/{id}", productHandler.Update).Methods("PATCH")
	apiRoutes.HandleFunc("/products/{id}", productHandler.Delete).Methods("DELETE")
	apiRoutes.HandleFunc("/products", productHandler.Create).Methods("POST")
	apiRoutes.HandleFunc("/products", productHandler.GetAll).Methods("GET")
}
