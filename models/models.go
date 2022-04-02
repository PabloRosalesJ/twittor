package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	Last      string             `bson:"last" json:"last,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Birthdate time.Time          `bson:"birthdate" json:"birthdate,omitempty"`
	Avatar    string             `bson:"avaar" json:"avaar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Bio       string             `bson:"bio" json:"bio,omitempty"`
	Web       string             `bson:"web" json:"web,omitempty"`
}
