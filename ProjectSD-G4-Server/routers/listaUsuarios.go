package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jayerdi20223/ProjectSD-G4/bd"
)

/*ListaUsuarios leo la lista de los usuarios */
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type") //capturar el parámetro type, q tipo de usuario queremos listar (aquellos usuarios q sigo o a aquellos q no sigo)
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search") //captura el término de búsqueda q puede o no venir

	pagTemp, err := strconv.Atoi(page) //convierte el alfabético en un entero
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp) //se convierte en entero 64 la página temporal

	result, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
