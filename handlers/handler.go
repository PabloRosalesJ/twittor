package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/PabloRosalesJ/twittor/middleware"
	"github.com/PabloRosalesJ/twittor/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/auth/register", middleware.CheckDB(routes.Register)).Methods("POST")
	router.HandleFunc("/auth/login", middleware.CheckDB(routes.Login)).Methods("POST")
	router.HandleFunc("/api/profile", middleware.CheckDB(middleware.ValidateJWT(routes.Login))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
