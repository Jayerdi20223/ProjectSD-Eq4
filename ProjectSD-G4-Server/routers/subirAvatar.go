package routers

import (
	"io" //se ve operaciones de input/ouput
	"net/http"
	"os"      //se ve funciones del SO
	"strings" //paquete q permite manejar string y realizar todo tipo de operaciones

	"github.com/Jayerdi20223/ProjectSD-G4/bd"
	"github.com/Jayerdi20223/ProjectSD-G4/models"
)

/*SubirAvatar sube el Avatar al servidor */
func SubirAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar") //viene en un formulario un archivo llamado avatar
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el avatar, reintente nuevamente"+err.Error(), 400)
		return
	}

	var extension = strings.Split(handler.Filename, ".")[1] //se va a capturar del archivo que vino la extensión (el elemento 1)
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666) //archivo, atributos q se da al archivo, permisos de lectura, escritura y ejecución q se da al archivo
	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	//una vez q abrió ese archivo en el disco. Es como que reservó espacio en el disco para ver si lo puede abrir o no
	_, err = io.Copy(f, file) //copia el archivo de FormFile en f
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + extension //se va a grabar en el campo Avatar el nombre del archivo q va a ser igual a IDUsuario.extensión
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la BD ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
