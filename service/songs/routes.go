package songs

import (
	"fmt"
	"net/http"
	"playr-server/types"
	"playr-server/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.SongStore
}

func NewHandler(store types.SongStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) SongsRoutes(router *mux.Router) {
	router.HandleFunc("/songs", h.handleGetSongs).Methods("GET")
	router.HandleFunc("/song/{id}", h.handleGetSongsById).Methods("GET")
	router.HandleFunc("/upload_song", h.handleUploadSong).Methods("POST")
	router.HandleFunc("/delete_song", h.handleDeleteSong).Methods("DELETE")
}

func (h *Handler) handleGetSongs(w http.ResponseWriter, r *http.Request) {
	songs, err := h.store.GetSongs()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "All Songs",
		"songs":   songs,
	})
}

func (h *Handler) handleGetSongsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songID, err := utils.ParseInt(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid song ID: %v", err))
		return
	}

	song, err := h.store.GetSongById(songID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"song": song,
	})
}

func (h *Handler) handleUploadSong(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("song")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	artist := r.FormValue("artist")
	name := r.FormValue("name")
	userID := r.FormValue("user_id")

	var userIDInt int
	_, err = fmt.Sscanf(userID, "%d", &userIDInt)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.store.UploadSong(file, header, artist, name, userIDInt)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to upload song: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Song uploaded successfully"))

}

func (h *Handler) handleDeleteSong(w http.ResponseWriter, r *http.Request) {
	var payload types.DeleteSongRequestPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	err := h.store.DeleteSong(payload.SongID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "Song Deleted")
}
