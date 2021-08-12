package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jayerdi20223/ProjectSD-G4/bd"
)

/*VerPerfil permite extraer los valores del Perfil */
func VerPerfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id") //se extrae del url el parámetro id
	if len(ID) < 1 {              //si no encuentra el id buscado
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest) //Devuelve requerimiento inválido
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil { //si no encontró el perfil del usuario
		http.Error(w, "Ocurrió un error al intentar buscar el registro "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json") //se va a devolver de tipo aplication/Json
	w.WriteHeader(http.StatusCreated)                  //devuelve un status 201 ya que no ha habido problemas
	json.NewEncoder(w).Encode(perfil)                  //devuelve en formato json
}
