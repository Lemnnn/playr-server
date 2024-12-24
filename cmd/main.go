package main

import (
	"log"
	"playr-server/cmd/api"
	"playr-server/pkg/database"
)

func main() {
	db, err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
