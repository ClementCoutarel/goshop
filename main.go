package main

import (
	"coutarel/goshop/config"
	"coutarel/goshop/database"
	"coutarel/goshop/router"
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

	_ = config.NewConfig(db, false)

	r := mux.NewRouter()

	userRepo := &database.UserSQLRepo{DB: db}
	productRepo := &database.ProductSQLRepo{DB: db}
	authRepo := &database.AuthSQLRepo{DB: db}

	userService := database.NewUserService(userRepo)
	productService := database.NewProductService(productRepo)
	authService := database.NewAuthService(authRepo)

	router.NewRouter(r, userService, productService, authService)

	server := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:3333",
	}

	log.Println("Server running")
	log.Fatal(server.ListenAndServe())
}
