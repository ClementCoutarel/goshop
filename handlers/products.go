package handlers

import (
	"coutarel/goshop/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	DB *sql.DB
}

// GetAll retrieves all the products from the database
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, name, description, price, quantity FROM products")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var products []models.Product

	for rows.Next() {
		var p models.Product
		rows.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Quantity)
		products = append(products, p)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetById retrieves a product by his id from the database
func (h *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var product models.Product
	err := h.DB.QueryRow("SELECT id, name, description, price, quantity FROM products WHERE id = ?", vars["id"]).Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}

}

// Create creates a new product from the database
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var exists int
	err = h.DB.QueryRow("SELECT id FROM products WHERE name = ?;", newProduct.Name).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = h.DB.Exec(
				"INSERT INTO products (name, description, price, quantity) VALUES(?,?,?,?);",
				newProduct.Name,
				newProduct.Description,
				newProduct.Price,
				newProduct.Quantity)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode("Product successfully created")
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if exists > 0 {
		http.Error(w, "Product already exists", http.StatusBadRequest)
	} else {

	}

}

// Update updates the infos of a product from the database
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	var updateProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	var exists int
	err = h.DB.QueryRow("SELECT id FROM products WHERE id = ?", updateProduct.Id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	_, err = h.DB.Exec("UPDATE products SET name = ?, description = ?, price = ?, quantity= ? WHERE id = ?;",
		updateProduct.Name,
		updateProduct.Description,
		updateProduct.Price,
		updateProduct.Quantity,
		updateProduct.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Product updated successfully")
}

// Delete deletes a product from the database
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Id format provided", http.StatusBadRequest)
		return
	}

	res, err := h.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Unable to delete the product", http.StatusInternalServerError)
		return
	}

	result, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Unable to identify the records affected", http.StatusInternalServerError)
		return
	}

	if result == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(200)
}
