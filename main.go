package main

import (
	"log"

	"github.com/SuperCodeOficial/twitt-max/bd"
	"github.com/SuperCodeOficial/twitt-max/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		bd.ChequeoConnection()
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

	// handlers.Saludo()
}
