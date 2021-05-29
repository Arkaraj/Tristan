package entity

type Project struct {
	_id         string  `json:"_id,omitempty" bson:"_id,omitempty"`
	title       string  `json:"title" bson:"title"`
	description string  `json:"description" bson:"description"`
	skills      []Skill `json:"skills" bson:"skills"`
	link        string  `json:"link" bson:"link"`
	startTime   string  `json:"startTime" bson:"startTime"`
	endTime     string  `json:"endTime" bson:"endTime"`
}
