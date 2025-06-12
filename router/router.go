package router

import (
	"coutarel/goshop/handlers"
	"database/sql"

	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, db *sql.DB) {

	userHandler := handlers.NewUserHandler(db)
	productHandler := handlers.NewProductHandler(db)

	apiRoutes := r.PathPrefix("/api").Subrouter()

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
