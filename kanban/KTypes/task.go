package KTypes

import "sync"

type Task struct {
	Mu          sync.Mutex    `json:"-" yaml:"-" bson:"-"`
	UUID        string        `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string        `json:"name" yaml:"name" bson:"name"`
	Type        string        `json:"type" yaml:"type" bson:"type"`
	Members     []Member      `json:"members,omitempty" yaml:"members" bson:"members"`
	Tags        []Tag         `json:"tags,omitempty" yaml:"tags" bson:"tags"`
	Priority    []Priority    `json:"priority,omitempty" yaml:"priority" bson:"priority"`
	Status      []Status      `json:"status,omitempty" yaml:"status" bson:"status"`
	DueDate     []DateDue     `json:"due_date,omitempty" yaml:"due_date" bson:"due_date"`
	Dates       []Date        `json:"dates,omitempty" yaml:"dates" bson:"dates"`
	Description []Description `json:"description,omitempty" yaml:"description" bson:"description"`
	Comments    []Comment     `json:"comments,omitempty" yaml:"comments" bson:"comments"`
	Attachments []Attachment  `json:"attachments,omitempty" yaml:"attachments" bson:"attachments"`
	CheckList   []CheckList   `json:"check_list,omitempty" yaml:"check_list" bson:"check_list"`
	Images      []Image       `json:"images,omitempty" yaml:"images" bson:"images"`
	Archived    []Archived    `json:"archived,omitempty" yaml:"archived" bson:"archived"`
	Activity    []Activity    `json:"activity,omitempty" yaml:"activity" bson:"activity"`
	Actions     []Actions     `json:"actions,omitempty" yaml:"actions" bson:"actions"`
}
