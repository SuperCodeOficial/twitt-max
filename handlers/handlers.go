package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

// Manejadores() Seteo mi puerto
func Manejadores() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServer(":"+PORT, handler))
}
