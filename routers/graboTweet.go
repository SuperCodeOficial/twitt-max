package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SuperCodeOficial/twitt-max/bd"
	"github.com/SuperCodeOficial/twitt-max/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
	}
	w.WriteHeader(http.StatusCreated)
}
