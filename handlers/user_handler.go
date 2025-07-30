package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/edigun-man/go-crud-api.git/models"
)

func parseParamId(w http.ResponseWriter, r *http.Request) (int, bool) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return 0, false
	}

	return id, true

}

// Get Users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Users)
}

// Get User By Id
func GetUser(w http.ResponseWriter, r *http.Request) {
	id, ok := parseParamId(w, r)
	if !ok {
		return
	}

	user := models.GetUserById(id)
	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// POST new User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	newUser := models.CreateUser(u.Name, u.Email)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)

}

// PUT User by Id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, ok := parseParamId(w, r)
	if !ok {
		return
	}

	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	updated := models.UpdateUser(id, u.Name, u.Email)

	if updated == nil {
		http.Error(w, "User not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

// DELETE User by Id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, ok := parseParamId(w, r)
	if !ok {
		return
	}

	deleted := models.DeleteUser(id)

	if !deleted {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
