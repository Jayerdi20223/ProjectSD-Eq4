package bd

import (
	"context"
	"time"

	"github.com/Jayerdi20223/ProjectSD-G4/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoTweets lee los tweets de un perfil */
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) { //trae (de golpe) todos los tweets (de la DB realizados) en un Slide
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()                          //se crea un objeto opciones
	opciones.SetLimit(20)                               //trae por paginación 20 documentos (tweets)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) //trae todos los documentos ordenados por la fecha en orden descendente
	opciones.SetSkip((pagina - 1) * 20)                 //De acuerdo al valor de página, se va ir paginando y salteando

	//cursor es un puntero, es como si fuera una tabla de DB donde se van a grabar los resultados y se va poder ir recorriendo y procesando de 1 a 1 a la vez
	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultados, false
	}

	for cursor.Next(context.TODO()) { //recorre c/documento guardado en el cursor

		var registro models.DevuelvoTweets //por c/iteración se va a crear un nuevo registro y se lo va a cargar
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro) //sirve para agregar en un slide un elemento
	}
	return resultados, true
}
