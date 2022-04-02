package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PabloRosalesJ/twittor/bd"
	"github.com/PabloRosalesJ/twittor/jwt"
	"github.com/PabloRosalesJ/twittor/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var model models.User

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, "User or password invalid "+err.Error(), 400)
		return
	}

	if len(model.Email) == 0 {
		http.Error(w, "The email is required", 400)
		return
	}

	document, exist := bd.Login(model.Email, model.Password)
	if !exist {
		http.Error(w, "User or password invalid", 400)
		return
	}

	fmt.Printf("%v", exist)
	fmt.Printf("%v", document)

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "We cant generate your token. "+err.Error(), 500)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

	expiredAt := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expiredAt,
	})
}
