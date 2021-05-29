package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Skill struct {
	Id             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Skill          string             `json:"skill" bson:"skill"`
	LevelOfMastery int                `json:"levelOfMastery" bson:"levelOfMastery"`
}
