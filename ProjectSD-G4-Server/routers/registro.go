package routers

import (
	"encoding/json" //paquete q permite trabajar con el formato Json
	"net/http"

	"github.com/Jayerdi20223/ProjectSD-G4/bd"
	"github.com/Jayerdi20223/ProjectSD-G4/models"
)

/*Registro es la funcion para crear en la BD el registro de usuario */
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario                      //creo un modelo json de Usuario
	err := json.NewDecoder(r.Body).Decode(&t) //el Body lo decodifica en el modelo. Todo lo q viene en el Body va terminar siendo una estructura de tipo Json

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true { //== true
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if status == false { //==false
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated) //devuelve el estatus del registro creado

}
