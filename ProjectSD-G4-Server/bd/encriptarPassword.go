package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la rutina que me permite encriptar la password recibida */
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo) //Slices de tipo byte que convierte al pass. Procesa la pass y lo encripta de manera eficiente
	return string(bytes), err
}

/*GenerateFromPassword implemeta un algoritmo de 2^costo, y esa es la cantidad de pasadas q se va a hacer al texto para encriptarlo
a mayor costo, más seguridad va a tener el password encriptado ya q va tener mayor cantidad de pasadas y más se demora.
2^8=256 pasadas, entonces el string (aunque sea 123456) q nosotros pasamos por parámetro lo va a encriptar una vez y a eso lo vuelve a encriptar y a ese resultado
lo vuelve a encriptar y así sucesivamente hasta las 256 veces lo q lo hace no hackeable.
Se recomienda:
Password de usuario común: 6, Password de SuperUsuario o Admin: 8, si supera esas cantidades puede llegar a relentizar mucho la ejecución del código,
lo que puede tardar varios segundos en encriptar una password*/
