package middlew

import (
	"net/http"

	"github.com/Jayerdi20223/ProjectSD-G4/routers"
)

/*ValidoJWT permite validar el JWT que nos viene en la petición */
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization")) //rutina q recibe el token por parámetro y devuelve si el token es válido o no
		if err != nil {
			http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest) //se ha realizado un petitorio inválido de acuerdo a la API
			return
		}
		next.ServeHTTP(w, r) //se pasa los objetos al siguiente eslabón de la cadena
	}
}
