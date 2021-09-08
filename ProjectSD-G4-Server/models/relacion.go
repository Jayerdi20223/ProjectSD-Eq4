package models

/*Relacion modelo para grabar la relacion de un usuario con otro */
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`                 //seteo q graba el id del usuario
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"` //seteo q graba el id del usuario que se esta siguiendo
}
