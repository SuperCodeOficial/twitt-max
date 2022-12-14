package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SuperCodeOficial/twitt-max/bd"
)

func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
	return
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro cmo entero mayor a 0 ! ", http.StatusBadRequest)
		return
	}
	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	if correcto == false {
		http.Error(w, "Error al leer los tweets ! ", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
