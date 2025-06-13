package database

import (
	"coutarel/goshop/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var product = models.Product{
	Name:        "ANdroid",
	Description: "An incredible phone",
	Price:       500,
	Quantity:    100,
}

var user = models.User{
	Name:     "Clement",
	Email:    "test@test.fr",
	Password: "123456",
	Role:     0,
}

func InitDb() *sql.DB {
	database, err := sql.Open("sqlite3", "./shop.db")
	if err != nil {
		log.Fatal(err)
	}

	query := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		role INT
    );
	
	CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
		description TEXT NOT NULL,
		price INT NOT NULL,
		QUANTITY INT
    );`
	_, err = database.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func SeedDb(db *sql.DB) {
	SeedUserTable(db)
	SeedProductTable(db)
}

func SeedProductTable(db *sql.DB) {
	checkQuery := `SELECT id FROM products WHERE name = ?;`

	var test int64
	err := db.QueryRow(checkQuery, product.Name).Scan(&test)
	if err != nil {
		if err == sql.ErrNoRows {
			insertQuery := `INSERT INTO products (name, description, price, quantity) VALUES(?,?,?,?);`
			_, err = db.Exec(insertQuery, product.Name, product.Description, product.Price, product.Quantity)
			if err != nil {
				log.Printf("Unable to seed the product table %s", err)
			} else {
				fmt.Printf("Product seeded successfully \n")
			}
		} else {
			fmt.Printf("Error accessing the db during seeding: %s \n", err.Error())
			return
		}
	} else {
		fmt.Printf("Product table already seeded \n")
	}

}

func SeedUserTable(db *sql.DB) {
	query := `SELECT id FROM users WHERE name = ? OR email = ?;`

	var test int64
	err := db.QueryRow(query, user.Name, user.Email).Scan(&test)
	if err != nil {
		if err == sql.ErrNoRows {
			query = `INSERT INTO users (name, email, password, role) VALUES(?,?,?,?);`

			_, err = db.Exec(query, user.Name, user.Email, user.Password, user.Role)
			if err != nil {
				fmt.Printf("Unable to seed the user table %s", err)
			} else {
				fmt.Printf("User table seeded successfully \n")
			}
		} else {
			fmt.Printf("Error accessing the user table during seeding %s \n", err.Error())
			return
		}
	} else {
		fmt.Printf("User table already seeded \n")
	}

}
