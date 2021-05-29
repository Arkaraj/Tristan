package entity

type Skill struct {
	id             string `json:id` // unique id, if i am using a db i.e
	sill           string `json:skill`
	levelOfMastery int    `json:levelOfMastery`
}
