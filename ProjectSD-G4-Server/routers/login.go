package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jayerdi20223/ProjectSD-G4/bd"
	"github.com/Jayerdi20223/ProjectSD-G4/jwt"
	"github.com/Jayerdi20223/ProjectSD-G4/models"
)

/*Login realiza el login */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe { //== false
		http.Error(w, "Usuario y/o Contraseña inválidos ", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento) //Recibe el documento y devuelve el token en modo string para q se pueda devolver al user con el http
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json") //se esta seteando
	w.WriteHeader(http.StatusCreated)                  //se esta devolviendo un status 200 o 201
	json.NewEncoder(w).Encode(resp)

	//Como se graba una Cookies desde el backend
	expirationTime := time.Now().Add(24 * time.Hour) //tiempo de expiración 24h desde que inicias en ese momento
	http.SetCookie(w, &http.Cookie{                  //función Cookie del http
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
