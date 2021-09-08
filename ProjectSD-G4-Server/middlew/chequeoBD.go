package middlew

import (
	"net/http"

	"github.com/Jayerdi20223/ProjectSD-G4/bd"
)

/*ChequeoBD es el middlew que me permite conocer el estado de la BD */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r) //le paso todos los valores q he recibido al próx. eslabón de la cadena
	}
}

//El middleware así como recibe el control q vino del navegador (http.HandlerFunc), también debe devolver la conexión http.HandlerFunc al siguiente paso.
