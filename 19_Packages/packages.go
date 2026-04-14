package packages

import (
	"fmt"
	"os/user"

	"github.com/Atharv-K-979/Go/19_Packages/auth"
)

func main() {

	auth.LoginWithCredientials("atharv", "password123")
	sessionToken := auth.GetSessionToken()
	fmt.Println("Session Token from main:", sessionToken)


	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current user: %s\n", user.Username)

}
