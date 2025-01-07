package auth

import (
	"fmt"
	"net/http"
	"playr-server/types"

	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
)

type Handler struct {
	store types.AuthStore
}

func NewHandler(store types.AuthStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) AuthRoutes(router *mux.Router) {
	fmt.Print(h)

	router.HandleFunc("/auth/{provider}/callback", h.handleGetAuthCallbackFunction).Methods("GET")
	router.HandleFunc("/logout/{provider}", h.handleLogoutFunction).Methods("GET")
	router.HandleFunc("/auth/{provider}", h.handleBeginAuthFunction).Methods("GET")
}

func (h *Handler) handleGetAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}

	newUser := &types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
	}

	err = h.store.CreateUser(newUser)
	if err != nil {
		fmt.Println("Error saving user to database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) handleLogoutFunction(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) handleBeginAuthFunction(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}
