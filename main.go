package main

import (
	"log"

	"github.com/joselimaico/twitt/bd"
	"github.com/joselimaico/twitt/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
