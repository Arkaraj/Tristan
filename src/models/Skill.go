package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Skill struct {
	_id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	sill           string             `json:"skill" bson:"skill"`
	levelOfMastery int                `json:"levelOfMastery" bson:"levelOfMastery"`
}
