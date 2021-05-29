package entity

type Project struct {
	id          string  `json:id` // unique id, if i am using a db i.e
	title       string  `json:title`
	description string  `json:description`
	skills      []Skill `json:skills`
	link        string  `json:link`
	startTime   string  `json:startTime`
	endTime     string  `json:endTime`
}
