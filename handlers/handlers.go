package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/SuperCodeOficial/twitt-max/middlew"
	"github.com/SuperCodeOficial/twitt-max/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores() Seteo mi puerto
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("GET")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leotweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("GET")
	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("POST")
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("POST")
	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/listaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leotweetsSeguidores", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

// ListenAndServer(":"+PORT, handler))
