package bd

import (
	"github.com/Jayerdi20223/ProjectSD-G4/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de login a la BD */
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado { // = false: si no se encontró el usuario que se intentó loguear
		return usu, false
	}

	passwordBytes := []byte(password)  //slide de tipo byte. Se esta grabando la password que vino del parámetro en esta función, no esta encriptada
	passwordBD := []byte(usu.Password) //Se esta grabando la password guardada en la DB que esta encriptada
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false //retorna usu ya que no lo encuentra
	}
	return usu, true
}
