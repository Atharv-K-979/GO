package auth

import "fmt"

func getSessionToken() string {
	return "session_token_123"
}

func LoginWithCredientials(username, password string) bool {
	fmt.Println("Login user with email and Password",username, password)
	token := getSessionToken()
	fmt.Println("Session Token:", token)
	return false
}
