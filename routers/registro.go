package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SuperCodeOficial/twitt-max/bd"
	"github.com/SuperCodeOficial/twitt-max/models"
)

// Registro
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos", 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El Email vino vacio, Llenelo", 400)
		return
	}
	if len(t.Password) < 7 {
		http.Error(w, "El password debe contener al menos 7 Caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Este usuario ya existe en la BD", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario", 400)
	}
	w.WriteHeader(http.StatusCreated)
}
