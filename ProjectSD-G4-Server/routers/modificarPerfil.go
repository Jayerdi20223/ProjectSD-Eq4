package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jayerdi20223/ProjectSD-G4/bd"
	"github.com/Jayerdi20223/ProjectSD-G4/models"
)

/*ModificarPerfil modifica el perfil de usuario */
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t) //se recibe en r.Body un Json y se decodifica en la posición de memoria de la variable t
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return //se aborta el método
	}

	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrión un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}

	if status == false { //No se ha encontrado el registro
		http.Error(w, "No se ha logrado modificar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
