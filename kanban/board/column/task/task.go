package task

import (
	"sharkedule/kanban/KTypes/actions"
	"sharkedule/kanban/KTypes/activity"
	"sharkedule/kanban/KTypes/archived"
	"sharkedule/kanban/KTypes/attachment"
	"sharkedule/kanban/KTypes/checklist"
	"sharkedule/kanban/KTypes/comment"
	"sharkedule/kanban/KTypes/date"
	"sharkedule/kanban/KTypes/datedue"
	"sharkedule/kanban/KTypes/description"
	"sharkedule/kanban/KTypes/image"
	"sharkedule/kanban/KTypes/member"
	"sharkedule/kanban/KTypes/priority"
	"sharkedule/kanban/KTypes/status"
	"sharkedule/kanban/KTypes/tag"
	"sync"
)

type Task struct {
	Mu          sync.Mutex                 `json:"-" yaml:"-" bson:"-"`
	UUID        string                     `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string                     `json:"name" yaml:"name" bson:"name"`
	Type        string                     `json:"type" yaml:"type" bson:"type"`
	Members     []*member.Member           `json:"members,omitempty" yaml:"members" bson:"members"`
	Tags        []*tag.Tag                 `json:"tags,omitempty" yaml:"tags" bson:"tags"`
	Priority    []*priority.Priority       `json:"priority,omitempty" yaml:"priority" bson:"priority"`
	Status      []*status.Status           `json:"status,omitempty" yaml:"status" bson:"status"`
	DueDate     []*datedue.DateDue         `json:"due_date,omitempty" yaml:"due_date" bson:"due_date"`
	Dates       []*date.Date               `json:"dates,omitempty" yaml:"dates" bson:"dates"`
	Description []*description.Description `json:"description,omitempty" yaml:"description" bson:"description"`
	Comments    []*comment.Comment         `json:"comments,omitempty" yaml:"comments" bson:"comments"`
	Attachments []*attachment.Attachment   `json:"attachments,omitempty" yaml:"attachments" bson:"attachments"`
	CheckList   []*checklist.CheckList     `json:"check_list,omitempty" yaml:"check_list" bson:"check_list"`
	Images      []*image.Image             `json:"images,omitempty" yaml:"images" bson:"images"`
	Archived    []*archived.Archived       `json:"archived,omitempty" yaml:"archived" bson:"archived"`
	Activity    []*activity.Activity       `json:"activity,omitempty" yaml:"activity" bson:"activity"`
	Actions     []*actions.Actions         `json:"actions,omitempty" yaml:"actions" bson:"actions"`
	Index       int                        `json:"index,omitempty" yaml:"index" bson:"index"`
	Board       string                     `json:"board,omitempty" yaml:"board" bson:"board"`
	Column      string                     `json:"column,omitempty" yaml:"column" bson:"column"`
}
