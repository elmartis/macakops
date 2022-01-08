package main

import (
	"github.com/elmartis/macakops/bd"
	"github.com/elmartis/macakops/handlers"
	"log"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
