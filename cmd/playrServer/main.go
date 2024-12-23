package main

import (
	"fmt"
	"net/http"

	"playr-server/internal/routes"
	"playr-server/pkg/supabase"
)

func main() {

	supabase.Init()
	router := routes.NewRouter()

	port := 8000
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n",addr)
	err:= http.ListenAndServe(addr,router)
	if err != nil {
		panic(err)
	}
}