package bd

import (
	"context"
	"time"

	"github.com/Jayerdi20223/ProjectSD-G4/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un email de parámetro y chequea si ya está en la BD */
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email} //Devuelve un map string, se setea con un formato Bson

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado) //devuelve un registro del contexto según la condición, convierte a Json el resultado
	ID := resultado.ID.Hex()                              //Trae el ObjectID del usuario q convierte un hex a formato string
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
