package bd

import (
	"context"
	"time"

	"github.com/Jayerdi20223/ProjectSD-G4/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro es la parada final con la BD para insertar los datos del usuario */
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //contexto con q vengo trabajando en la DB, 15 seg.
	defer cancel()                                                           //cancela el context.WithTimeout por medio del objeto cancel

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password) //debido a que no se pueden grabar password ingresados en formato texto en la DB por medidas de seguridad.

	result, err := col.InsertOne(ctx, u) //Inserta un registro del contexto del modelo
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
