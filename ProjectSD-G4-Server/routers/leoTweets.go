package routers

import (
	"encoding/json"
	"net/http"
	"strconv" //paquete para hacer conversión de datos

	"github.com/Jayerdi20223/ProjectSD-G4/bd"
)

/*LeoTweets Leo los tweets */
func LeoTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) //conversión de un string a un entero
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagina) //se va a convertir una pag de tipo int a tipo int64
	respuesta, correcto := bd.LeoTweets(ID, pag)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json") //seteamos el header diciendole que va a ser de ese tipo
	w.WriteHeader(http.StatusCreated)                  //se hace una escritura en el header diciendo q fue satisfactorio
	json.NewEncoder(w).Encode(respuesta)
}
