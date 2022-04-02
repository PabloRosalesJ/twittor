package routes

import (
	"encoding/json"
	"net/http"

	"github.com/PabloRosalesJ/twittor/bd"
	"github.com/PabloRosalesJ/twittor/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Bad request. "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "The password must contain at least 6 characters", 400)
		return
	}

	_, founded, _ := bd.ExistUser(t.Email)
	if founded {
		http.Error(w, "The email already taken", 400)
		return
	}

	_, stored, err := bd.StoreUser(t)
	if err != nil {
		http.Error(w, "Error to creae user. "+err.Error(), 500)
		return
	}
	if !stored {
		http.Error(w, "Can't create user. ", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
