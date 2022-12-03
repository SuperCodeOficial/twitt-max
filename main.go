package main

import (
	"log"

	"github.com/SuperCodeOficial/twitt-max/bd"
	"github.com/SuperCodeOficial/twitt-max/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
