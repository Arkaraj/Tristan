package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	_id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	name        string             `json:"name" bson:"name"`
	description string             `json:"description" bson:"name"`
	tag         string             `json:"tag" bson:"tag"`
	skills      []Skill            `json:"skills" bson:"skills"`
	projects    []Project          `json:"projects" bson:"projects"`
}
