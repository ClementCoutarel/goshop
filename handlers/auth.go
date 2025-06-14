package handlers

import (
	"coutarel/goshop/database"
	"coutarel/goshop/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	AuthService database.AuthService
}

func NewAuthHandler(authService database.AuthService) AuthHandler {
	return AuthHandler{
		AuthService: authService,
	}
}

// Signin signs in the user if the credentials or validated
func (h *AuthHandler) Signin(w http.ResponseWriter, r *http.Request) {
	var credentials models.UserToAuthDTO
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Unable to check ther request body provided", http.StatusInternalServerError)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var registeredUser models.User
	err = h.DB.QueryRow("SELECT id,name,email,password,role FROM users WHERE email = ?", &credentials.Email).Scan(
		&registeredUser.Id,
		&registeredUser.Name,
		&registeredUser.Email,
		&registeredUser.Password,
		&registeredUser.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if credentials.Password != registeredUser.Password {
		http.Error(w, "Password don't match", http.StatusUnauthorized)
		return
	}

	if credentials.Email != registeredUser.Email {
		http.Error(w, "Email don't match", http.StatusUnauthorized)
		return
	}

	token, err := CreateToken(registeredUser.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	loggedUser := models.LoggedUserDTO{
		Id:    registeredUser.Id,
		Name:  registeredUser.Name,
		Email: registeredUser.Email,
		Role:  registeredUser.Role,
		Token: token,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(loggedUser)

}

// Register rergisters the user with the provided body
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var newUser models.UserCreateDTO
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid body provided", http.StatusBadRequest)
		return
	}

	var exists int
	err = h.DB.QueryRow("SELECT COUNT(*) FROM USERS WHERE name = ? OR email = ?", &newUser.Name, &newUser.Email).Scan(&exists)
	if err != nil {
		http.Error(w, "An error occured while accessing the database", http.StatusInternalServerError)
		return
	}

	if exists > 0 {
		http.Error(w, "User already registered", http.StatusConflict)
		return
	}

	res, err := h.DB.Exec("INSERT INTO users (name, email, password,role) VALUES(?,?,?,?)", &newUser.Name, &newUser.Email, &newUser.Password, &newUser.Role)
	if err != nil {
		http.Error(w, "An error occured during registration", http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Unable to check the user inserted", http.StatusInternalServerError)
		return
	}

	createdUser := models.UserGetDTO{
		Id:    int(id),
		Name:  newUser.Name,
		Email: newUser.Email,
		Role:  newUser.Role,
	}

	w.Header().Set("Content-type", "apllication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
