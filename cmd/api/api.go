package api

import (
	"database/sql"
	"log"
	"net/http"
	"playr-server/service/auth"
	"playr-server/service/liked_songs"
	"playr-server/service/songs"
	"playr-server/service/users"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := users.NewHandler()
	userHandler.UserRoutes(subrouter)

	authStore := auth.NewStore(s.db)
	authHandler := auth.NewHandler(authStore)
	authHandler.AuthRoutes(subrouter)

	songsStore := songs.NewStore(s.db)
	songHandler := songs.NewHandler(songsStore)
	songHandler.SongsRoutes(subrouter)

	likedSongsStore := liked_songs.NewStore(s.db)
	likeHandler := liked_songs.NewHandler(likedSongsStore)
	likeHandler.LikedRoutes(subrouter)

	log.Println("Listening", s.addr)
	return http.ListenAndServe(s.addr, router)
}
