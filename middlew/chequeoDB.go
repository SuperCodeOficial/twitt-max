package middlew

import (
	"net/http"

	"github.com/SuperCodeOficial/twitt-max/bd"
)

// ChequeoBD
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Se perdio la conexion con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
