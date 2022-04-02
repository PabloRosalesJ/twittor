package bd

import (
	"github.com/PabloRosalesJ/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.User, bool) {
	user, userFounded, _ := ExistUser(email)

	if !userFounded {
		return user, false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, false
	}
	return user, true

}
