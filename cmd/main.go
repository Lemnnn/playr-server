package main

import (
	"log"
	"playr-server/cmd/api"
	"playr-server/pkg/database"
	"playr-server/service/auth"
)

func main() {

	db, err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	auth.NewAuth()
	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
