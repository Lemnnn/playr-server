package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"playr-server/internal/users"
)

func NewRouter() http.Handler{
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/signup", signUpHandler)

	return mux
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := users.Register(credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Registration failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}