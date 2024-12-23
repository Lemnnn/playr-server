package users

import (
	"fmt"
	"playr-server/pkg/supabase"
)

func Register(email, password string) error {
	client := supabase.Client

	user, err := client.SignInWithEmailPassword(email, password)

	if err != nil {
		return err
	}
	
	fmt.Println(user)
	return nil
}