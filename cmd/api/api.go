package api

import (
	"database/sql"
	"log"
	"net/http"
	"playr-server/service/auth"
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

	log.Println("Listening", s.addr)
	return http.ListenAndServe(s.addr, router)
}
