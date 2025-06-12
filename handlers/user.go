package handlers

import (
	"coutarel/goshop/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	DB *sql.DB
}

// Update updates a user information
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id provided", http.StatusBadRequest)
		return
	}

	var newUser models.User
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	var test int
	err = h.DB.QueryRow("SELECT id FROM users WHERE id = ?;", id).Scan(&test)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, "Unable to update the user", http.StatusInternalServerError)
			return
		}
	}

	_, err = h.DB.Exec("UPDATE users SET name = ?, email= ?, password = ?,role = ? WHERE id = ?",
		newUser.Name,
		newUser.Email,
		newUser.Password,
		newUser.Role,
		id,
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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID provided", http.StatusBadRequest)
		return
	}

	res, err := h.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Unable to find the user affected", http.StatusInternalServerError)
		return
	}

	if result == 0 {
		http.Error(w, "UNable to find the user for deletion", http.StatusNotFound)
		return
	}

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

	var exists string
	err = h.DB.QueryRow("SELECT name FROM users WHERE name = ? OR email = ?", newUser.Name, newUser.Email).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
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
			return
		} else {
			http.Error(w, "User already registered", http.StatusInternalServerError)
			return
		}
	}

}

// GetById retreves all the users from the database
func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID provided", http.StatusBadRequest)
		return
	}

	var user models.UserDTO
	err = h.DB.QueryRow("SELECT id, name, email, role FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Email, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
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
		rows.Scan(&u.Id, &u.Name, &u.Email, &u.Role)
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}
