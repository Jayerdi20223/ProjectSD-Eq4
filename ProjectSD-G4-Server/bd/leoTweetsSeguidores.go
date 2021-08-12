package bd

import (
	"context"
	"time"

	"github.com/Jayerdi20223/ProjectSD-G4/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTweetsSeguidores lee los tweets de mis seguidores */
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20 //se va a paginar de a 20 resultados

	condiciones := make([]bson.M, 0)                                             //se fuerza q el slide tenga 0 elementos
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}}) //match busca el usuarioid de la "tabla" relación para así filtrar ese id y traer todos sus registros
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{ //lookup permite unir dos tablas. Se ponen los cuatro parámetros necesarios para unir tablas en mongo
			"from":         "tweet",             //en from se pone con q tabla se quiere unir la tabla relación, en este caso sería la tabla tweet
			"localField":   "usuariorelacionid", //es el campo local por el cual se va a unir las tablas
			"foreignField": "userid",            //es el campo de la tabla tweet donde esta el usuario id (es el que va a ver los tweets de seguidores)
			"as":           "tweet",             //es el alias de como queremos llamar la tabla tweet, en este caso se llamará tweet
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})                //unwind permite q todos los documentos vengan iguales, es decir, van a venir documentos con información repetida (del q realizó el tweet) en todos los tweets que haga
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}}) //datos ordenados por el campo fecha y de orden descendente (último primero, el más antiguo al final)
	condiciones = append(condiciones, bson.M{"$skip": skip})                      //paginar, es decir, sortear o saltar 20 registros en la búsqueda
	condiciones = append(condiciones, bson.M{"$limit": 20})                       //limit para leer 20 registros

	cursor, err := col.Aggregate(ctx, condiciones) //se va a ejecutar directamente en la DB con esa condiciones y se va a crear un cursor
	var result []models.DevuelvoTweetsSeguidores
	err = cursor.All(ctx, &result) //se ejecuta todo el cursor (arma todo el documento q se tiene q enviar) y el resultado se decodifica conforme al modelo
	if err != nil {
		return result, false
	}
	return result, true
}
