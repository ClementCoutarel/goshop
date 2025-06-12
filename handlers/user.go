package handlers

import (
	"coutarel/goshop/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	DB *sql.DB
}

// Update updates a user information
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	rows, err := h.DB.Query("SELECT id,name,email FROM users WHERE id = ?", vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.Id, &u.Name)
		users = append(users, u)
	}

	if len(users) == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	_, err = h.DB.Exec("UPDATE users SET name = ?, email= ?, password = ?,role = ? WHERE id = ?",
		newUser.Name,
		newUser.Email,
		newUser.Password,
		newUser.Role,
		vars["id"],
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(newUser)

}

// Delete deletes a user from the database
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows, err := h.DB.Query("SELECT id,name,email FROM users WHERE id = ?;", vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.Id, &u.Name)
		users = append(users, u)
	}

	if len(users) == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	res, err := h.DB.Exec("DELETE FROM users WHERE id = ?", vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(res)

	w.WriteHeader(204)
}

// Create creates a new user in the database
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	rows, err := h.DB.Query("SELECT id,name,email FROM users WHERE name = ? OR email = ?", newUser.Name, newUser.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.Id, &u.Name)
		users = append(users, u)
	}

	if len(users) > 0 {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	res, err := h.DB.Exec("INSERT INTO users (name, email, password) VALUES(?,?,?);",
		newUser.Name,
		newUser.Email,
		newUser.Password,
		0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	newUser.Id = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUser)

}

// GetById retreves all the users from the database
func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rows, err := h.DB.Query("SELECT id,name,email, role FROM users WHERE id = ?", vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	var users []models.UserDTO
	for rows.Next() {
		var u models.UserDTO
		rows.Scan(&u.Id, &u.Name, &u.Email, &u.Role)
		users = append(users, u)
	}

	if len(users) <= 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// GetAll retrieves all users from the database
func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, name, email, role FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.UserDTO
	for rows.Next() {
		var u models.UserDTO
		rows.Scan(&u.Id, &u.Name, &u.Email, *&u.Role)
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}
