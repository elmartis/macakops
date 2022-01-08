package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net"
	"os"
)

/*manejadores de seteo mi puerto el handlers y pongo a escuchar al server*/
func Manejadores() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.listernAndServe(":"+PORT, handler))
}
