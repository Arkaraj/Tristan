package entity

type User struct {
	name        string    `json:name`
	description string    `json:description`
	tag         string    `json:tag`
	skills      []Skill   `json:skills`
	projects    []Project `json:projects`
}
