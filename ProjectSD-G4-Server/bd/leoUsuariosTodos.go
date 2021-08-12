package bd

import (
	"context"
	"time"

	"github.com/Jayerdi20223/ProjectSD-G4/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoUsuariosTodos Lee los usuarios registrados en el sistema, si se recibe "R" en quienes
  trae solo los que se relacionan conmigo */
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.Usuario //se va a enviar un slide de usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search}, //?i: No se fija si es mayúscula o minúscula. Se va a buscar dentro del string sin importar si es mayúscula o minúscula
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) { //se va a recorrer el cursor. Next permite avanzar al siguiente registro
		var s models.Usuario
		err := cur.Decode(&s) //se va grabando el cursor en modelo usuario
		if err != nil {
			return results, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex() //se extrae el ObjectID

		incluir = false

		encontrado, err = ConsultoRelacion(r)

		if tipo == "new" && encontrado == false { //lista de usuarios al q no se siguen
			incluir = true
		}
		if tipo == "follow" && encontrado == true { //listado de los usuarios a los q sigo
			incluir = true
		}

		if r.UsuarioRelacionID == ID { //si el usuario y el usuario al que se sigue son lo mismo, entonces no incluirlo
			incluir = false
		}

		if incluir == true { //se blanquea todos los campos q no me interesa incluir en el listado
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s) //se va a grabar en el puntero de memoria, va a extraer los datos y se va a adicionar al slide
		}
	}

	err = cur.Err()
	if err != nil {
		return results, false
	}
	cur.Close(ctx)
	return results, true
}

/* ID: es el q esta leyendo a los demás usuarios;
page: se va a listar usuarios, pero se tiene que paginar de a 20 resultados;
search: nos va a permitir inclusive filtrar por una palabra, por un término;
tipo: vamos a indicar q tipo de búsqueda vamos a hacer, si vamos a listar todos los usuarios,
si vamos a listar solamente los que nos siguen a nosotros o con los cuales nosotros seguimos */
