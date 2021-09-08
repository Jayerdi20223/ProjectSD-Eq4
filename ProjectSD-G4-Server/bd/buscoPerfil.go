package bd

import (
	"context"
	"time"

	"github.com/Jayerdi20223/ProjectSD-G4/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil busca un perfil en la BD */
func BuscoPerfil(ID string) (models.Usuario, error) { //ID como parámetro de entrada al buscarlo por GET en la url
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var perfil models.Usuario                 //se va a grabar todo el perfil del usuario q vamos a consultar
	objID, _ := primitive.ObjectIDFromHex(ID) //convierte el string en un objeto de tipo ID

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil) //Devolución de la información en perfil
	perfil.Password = ""                               //para omitir el password o te devuelva vacío ya que no se quiere mostrar el password del usuario
	if err != nil {
		return perfil, err
	}
	return perfil, nil
}
