package liked_songs

import (
	"fmt"
	"net/http"
	"playr-server/types"
	"playr-server/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.LikedSongsStore
}

func NewHandler(store types.LikedSongsStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) LikedRoutes(router *mux.Router) {
	fmt.Print(h)

	router.HandleFunc("/addLike", h.handleAddToLikedSongs).Methods("POST")
	router.HandleFunc("/myLikes", h.handleGetUserLikedSongs).Methods("GET")
	router.HandleFunc("/removeLike", h.handleRemoveFromLikedSongs).Methods("DELETE")
}

func (h *Handler) handleAddToLikedSongs(w http.ResponseWriter, r *http.Request) {
	var payload types.AddLikeRequestPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	err := h.store.AddToLikedSongs(payload.UserID, payload.SongID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "Added to Liked Songs")
}

func (h *Handler) handleGetUserLikedSongs(w http.ResponseWriter, r *http.Request) {
	var payload types.GetUserLikesRequestPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	likedSongs, err := h.store.GetUserLikedSongs(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "My Liked Songs",
		"songs":   likedSongs,
	})
}

func (h *Handler) handleRemoveFromLikedSongs(w http.ResponseWriter, r *http.Request) {
	var payload types.RemoveLikeRequestPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	err := h.store.RemoveFromLikedSongs(payload.LikedID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "Removed to Liked Songs")
}
