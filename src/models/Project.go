package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	Id          primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string               `json:"title" bson:"title"`
	Description string               `json:"description" bson:"description"`
	Skills      []primitive.ObjectID `json:"skills" bson:"skills"`
	Link        string               `json:"link" bson:"link"`
	StartTime   time.Time            `json:"startTime" bson:"startTime" time_format:"2006-01-02" time_utc:"1"`
	EndTime     time.Time            `json:"endTime" bson:"endTime" time_format:"2006-01-02" time_utc:"1"`
}

// Example startTime: 2030-04-16, endTime: 2030-04-17
