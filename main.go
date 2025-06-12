package main

import (
	"coutarel/goshop/config"
	"coutarel/goshop/database"
	"coutarel/goshop/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	db := database.InitDb()
	database.SeedDb(db)
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Unable to reach the database \n", err)
	}

	config := &config.Config{
		DB: db,
	}

	userHandler := handlers.UserHandler{
		DB: config.DB,
	}

	productHandler := handlers.ProductHandler{
		DB: config.DB,
	}

	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")
	r.HandleFunc("/users/{id}", userHandler.Update).Methods("PATCH")
	r.HandleFunc("/users/{id}", userHandler.GetById).Methods("GET")
	r.HandleFunc("/users", userHandler.GetAll).Methods("GET")
	r.HandleFunc("/users", userHandler.Create).Methods("POST")

	r.HandleFunc("/products/{id}", productHandler.GetById).Methods("GET")
	r.HandleFunc("/products/{id}", productHandler.Update).Methods("PATCH")
	r.HandleFunc("/products/{id}", productHandler.Delete).Methods("DELETE")
	r.HandleFunc("/products", productHandler.Create).Methods("POST")
	r.HandleFunc("/products", productHandler.GetAll).Methods("GET")

	server := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
	}

	log.Println("Server running")
	log.Fatal(server.ListenAndServe())
}
