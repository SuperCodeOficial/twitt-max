package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoTweet struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omtempty"`
	UserID  string             `bson:"usuarioid" json:"usuarioid,omtempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omtempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omtempty"`
}
