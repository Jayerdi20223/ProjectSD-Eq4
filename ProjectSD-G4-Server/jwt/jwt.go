package jwt

import (
	"time"

	"github.com/Jayerdi20223/ProjectSD-G4/models"
	jwt "github.com/dgrijalva/jwt-go" //se crea el alias "jwt" para no tener q mencionar el nombre "jwt-go", a partir de ahí dentro del paquete lo vamos a llamar jwt directamente
)

/*GeneroJWT genera el encriptado con JWT */
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("MastersdelDesarrollo_grupodeFacebook") //clave privada en formato byte

	payload := jwt.MapClaims{ //Es la parte de los datos. Se crea la lista de privilegios a guardar en el payload
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(), //el formato Unix convierte la fecha en formato long
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) //Se esta eligiendo el algoritmo para encriptar y hacer todo el cálculo, se pasa el payload
	tokenStr, err := token.SignedString(miClave)                //se establece el string de firma
	if err != nil {
		return tokenStr, err //devuelve un string vacío
	}
	return tokenStr, nil
}
