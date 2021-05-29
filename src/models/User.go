package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"description" bson:"name"`
	Tag         string               `json:"tag" bson:"tag"`
	Skills      []primitive.ObjectID `json:"skills" bson:"skills"`
	Projects    []primitive.ObjectID `json:"projects" bson:"projects"`
}
