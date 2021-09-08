package bd

import (
	"context"
	"time"

	"github.com/Jayerdi20223/ProjectSD-G4/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificoRegistro permite modificar el perfil del usuario */
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{}) //todo el mapa va a ser de tipo interface. Se está creando una interface vacía
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre //se esta creando una clave dentro de ese mapa para el campo nombre y se esta grabando un valor u.Nombre
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	registro["fechaNacimiento"] = u.FechaNacimiento //se fuerza el valor del que viene del modelo de datos
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	updtString := bson.M{ //se arma el registro de actualización
		"$set": registro, //esto se va a actualizar
	}

	objID, _ := primitive.ObjectIDFromHex(ID)     //convierte el string ID a un formato objectID
	filtro := bson.M{"_id": bson.M{"$eq": objID}} //condición _id = objID

	_, err := col.UpdateOne(ctx, filtro, updtString) //solo se va a encontrar un registro o documento a actualizar
	if err != nil {
		return false, err
	}

	return true, nil
}
