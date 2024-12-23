package supabase

import (
	"fmt"

	"github.com/supabase-community/supabase-go"
)

var Client *supabase.Client

func Init() {
    supabaseURL := "https://acyldzrfohtwldwufzuo.supabase.co" 
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImFjeWxkenJmb2h0d2xkd3VmenVvIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MzQ3Nzc5MDcsImV4cCI6MjA1MDM1MzkwN30.6GuDK9DOsXiJhJva8UydqVAsatp_LLazAGh-FR2PETY"

	client, err := supabase.NewClient(supabaseURL, supabaseKey, nil)
	if err != nil {
	  fmt.Println("cannot initalize client", err)
	}

	Client = client
}