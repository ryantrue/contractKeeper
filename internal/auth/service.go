package auth

import (
	"github.com/ryantrue/contractKeeper.git/internal/data"
)

func AuthenticateUser(username, password string) bool {
	return data.ValidateUserCredentials(username, password)
}
