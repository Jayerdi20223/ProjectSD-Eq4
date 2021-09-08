package models

import "time"

/*GraboTweet es el formato o estructura que tendrá nuestro Tweet en la BD */
type GraboTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"` //userid: es el nombre q tendrá en DB
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
