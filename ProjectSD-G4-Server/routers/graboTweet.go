package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jayerdi20223/ProjectSD-G4/bd"
	"github.com/Jayerdi20223/ProjectSD-G4/models"
)

/*GraboTweet permite grabar el tweet en la base de datos */
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje) //se recibe el Json en el body y se ha copiado y decodificado dentro de mensaje

	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return //se aborta el método
	}

	registro := models.GraboTweet{ //en la variable se guarda toda la información
		UserID:  IDUsuario,       //en el campo UserID se va a grabar IDUsuario
		Mensaje: mensaje.Mensaje, //en el campo se graba 'mensaje' que es el body q se ha decodificado y adento tiene un campo llamado 'Mensaje' q es lo q viene en el Json del body
		Fecha:   time.Now(),      //esa función devuelve la fecha de hoy con h,m,s.
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
