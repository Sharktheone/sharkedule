package kanban

import "sync"

type Board struct {
	Mu          sync.Mutex   `json:"-" yaml:"-" bson:"-"`
	UUID        string       `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string       `json:"name" yaml:"name" bson:"name"`
	Description Description  `json:"description,omitempty" yaml:"description" bson:"description"`
	Members     []Member     `json:"members,omitempty" yaml:"members" bson:"members"`
	Tags        []Tag        `json:"tags,omitempty" yaml:"tags" bson:"tags"`
	Priority    Priority     `json:"priority,omitempty" yaml:"priority" bson:"priority"`
	Status      Status       `json:"status,omitempty" yaml:"status" bson:"status"`
	DueDate     DateDue      `json:"due_date,omitempty" yaml:"due_date" bson:"due_date"`
	Dates       []Date       `json:"dates,omitempty" yaml:"dates" bson:"dates"`
	Comments    []Comment    `json:"comments,omitempty" yaml:"comments" bson:"comments"`
	Attachments []Attachment `json:"attachments,omitempty" yaml:"attachments" bson:"attachments"`
	CheckList   []CheckList  `json:"check_list,omitempty" yaml:"check_list" bson:"check_list"`
	Images      []Image      `json:"images,omitempty" yaml:"images" bson:"images"`
	Archived    Archived     `json:"archived,omitempty" yaml:"archived" bson:"archived"`
	Activity    []Activity   `json:"activity,omitempty" yaml:"activity" bson:"activity"`
	Actions     []Actions    `json:"actions,omitempty" yaml:"actions" bson:"actions"`
	Columns     []Column     `json:"columns,omitempty" yaml:"columns" bson:"columns"`
}

type Column struct {
	UUID        string           `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string           `json:"name" yaml:"name" bson:"name"`
	Description Description      `json:"description,omitempty" yaml:"description" bson:"description"`
	Tasks       []KanbanTaskType `json:"tasks,omitempty" yaml:"tasks" bson:"tasks"`
}

type KanbanTaskType struct {
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

type Description struct {
	Description string `json:"description,omitempty" yaml:"description" bson:"description"`
}

type Priority struct {
	UUID     string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Priority string `json:"priority" yaml:"priority" bson:"priority"`
}

type Status struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
}

type DateDue struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Date string `json:"date" yaml:"date" bson:"date"`
}

type Archived struct {
	Archived bool   `json:"archived" yaml:"archived" bson:"archived"`
	Date     string `json:"date" yaml:"date" bson:"date"`
	User     string `json:"user" yaml:"user" bson:"user"`
}

type Member struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
}

type Tag struct {
	UUID  string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name  string `json:"name" yaml:"name" bson:"name"`
	Color string `json:"color" yaml:"color" bson:"color"`
	Icon  string `json:"icon,omitempty" yaml:"icon" bson:"icon"`
}

type Date struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name string `json:"name" yaml:"name" bson:"name"`
	Date string `json:"date" yaml:"date" bson:"date"`
}

type Attachment struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name string `json:"name" yaml:"name" bson:"name"`
	Type string `json:"type" yaml:"type" bson:"type"`
	Size string `json:"size" yaml:"size" bson:"size"`
	Date string `json:"date" yaml:"date" bson:"date"`
	User string `json:"user" yaml:"user" bson:"user"`
}

type Image struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name string `json:"name" yaml:"name" bson:"name"`
	Type string `json:"type" yaml:"type" bson:"type"`
	Size string `json:"size" yaml:"size" bson:"size"`
	Date string `json:"date" yaml:"date" bson:"date"`
	User string `json:"user" yaml:"user" bson:"user"`
	URL  string `json:"url" yaml:"url" bson:"url"`
}

type Comment struct {
	UUID    string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Message string `json:"message" yaml:"message" bson:"message"`
	Date    string `json:"date" yaml:"date" bson:"date"`
	User    string `json:"user" yaml:"user" bson:"user"`
}

type CheckList struct {
	UUID  string                    `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name  string                    `json:"name" yaml:"name" bson:"name"`
	Items []KanbanCheckListItemType `json:"items,omitempty" yaml:"items" bson:"items"`
}

type KanbanCheckListItemType struct {
	UUID    string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name    string `json:"name" yaml:"name" bson:"name"`
	Checked bool   `json:"checked" yaml:"checked" bson:"checked"`
}

type Actions struct {
	UUID   string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name   string `json:"name" yaml:"name" bson:"name"`
	Icon   string `json:"icon" yaml:"icon" bson:"icon"`
	Color  string `json:"color" yaml:"color" bson:"color"`
	Type   string `json:"type" yaml:"type" bson:"type"`
	Action string `json:"action" yaml:"action" bson:"action"`
}

type Activity struct {
	UUID    string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Message string `json:"message" yaml:"message" bson:"message"`
	Date    string `json:"date" yaml:"date" bson:"date"`
	User    string `json:"user" yaml:"user" bson:"user"`
}
