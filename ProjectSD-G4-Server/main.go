package main

import (
	"log"

	"github.com/Jayerdi20223/ProjectSD-G4/bd"
	"github.com/Jayerdi20223/ProjectSD-G4/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
