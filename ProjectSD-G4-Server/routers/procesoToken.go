package routers

import (
	"errors"
	"strings"

	"github.com/Jayerdi20223/ProjectSD-G4/bd"
	"github.com/Jayerdi20223/ProjectSD-G4/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email es el valor de Email usado en todos los EndPoints */
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints */
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer") //logra q el splitToken se convierta en un vector el cual en el elemento 0 se va a tener Bearer y en el elemento 1 tendré el String pegado
	if len(splitToken) != 2 {                 //validación de q el token tenga dos elementos
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1]) //conversión del token, la función quita los espacios que haya

	// Validación misma del token
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) { //interface{} convierte []byte("MastersdelDesarrollo_grupodeFacebook") en un objeto Json
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true { //== true
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido") //errors: para string
	}
	return claims, false, string(""), err //Devuelve err: porque es un objeto de tipo error
}
