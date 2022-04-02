package main

import (
	"log"

	"github.com/PabloRosalesJ/twittor/bd"
	"github.com/PabloRosalesJ/twittor/handlers"
)

func main() {
	if !bd.CheckConnetion() {
		log.Fatal("No connection")
	}
	handlers.Handlers()
}
