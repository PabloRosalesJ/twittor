package jwt

import (
	"time"

	"github.com/PabloRosalesJ/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Secret = []byte("DBBalLzY1zc11N4SC6MD")

func GenerateJWT(m models.User) (string, error) {

	payload := jwt.MapClaims{
		"email":     m.Email,
		"name":      m.Name,
		"last":      m.Last,
		"birthdate": m.Birthdate,
		"web":       m.Web,
		"_id":       m.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(Secret)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, err
}
