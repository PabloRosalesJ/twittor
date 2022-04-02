package middleware

import (
	"net/http"

	"github.com/PabloRosalesJ/twittor/bd"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !bd.CheckConnetion() {
			http.Error(w, "Connection lossed to DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
