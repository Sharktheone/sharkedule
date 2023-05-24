package kanbanboardTypes

import "sync"

type KanbanBoard struct {
	Mu          sync.Mutex             `json:"-" yaml:"-" bson:"-"`
	UUID        string                 `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string                 `json:"name" yaml:"name" bson:"name"`
	Description KanbanDescriptionType  `json:"description,omitempty" yaml:"description" bson:"description"`
	Members     []KanbanMemberType     `json:"members,omitempty" yaml:"members" bson:"members"`
	Tags        []KanbanTagType        `json:"tags,omitempty" yaml:"tags" bson:"tags"`
	Priority    KanbanPriorityType     `json:"priority,omitempty" yaml:"priority" bson:"priority"`
	Status      KanbanStatusType       `json:"status,omitempty" yaml:"status" bson:"status"`
	DueDate     KanbanDateDueType      `json:"due_date,omitempty" yaml:"due_date" bson:"due_date"`
	Dates       []KanbanDateType       `json:"dates,omitempty" yaml:"dates" bson:"dates"`
	Comments    []KanbanCommentType    `json:"comments,omitempty" yaml:"comments" bson:"comments"`
	Attachments []KanbanAttachmentType `json:"attachments,omitempty" yaml:"attachments" bson:"attachments"`
	CheckList   []KanbanCheckListType  `json:"check_list,omitempty" yaml:"check_list" bson:"check_list"`
	Images      []KanbanImageType      `json:"images,omitempty" yaml:"images" bson:"images"`
	Archived    KanbanArchivedType     `json:"archived,omitempty" yaml:"archived" bson:"archived"`
	Activity    []KanbanActivityType   `json:"activity,omitempty" yaml:"activity" bson:"activity"`
	Actions     []KanbanActionType     `json:"actions,omitempty" yaml:"actions" bson:"actions"`
	Columns     []KanbanColumnType     `json:"columns,omitempty" yaml:"columns" bson:"columns"`
}

type KanbanColumnType struct {
	UUID        string                `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string                `json:"name" yaml:"name" bson:"name"`
	Description KanbanDescriptionType `json:"description,omitempty" yaml:"description" bson:"description"`
	Tasks       []KanbanTaskType      `json:"tasks,omitempty" yaml:"tasks" bson:"tasks"`
}

type KanbanTaskType struct {
	UUID        string                  `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string                  `json:"name" yaml:"name" bson:"name"`
	Type        string                  `json:"type" yaml:"type" bson:"type"`
	Members     []KanbanMemberType      `json:"members,omitempty" yaml:"members" bson:"members"`
	Tags        []KanbanTagType         `json:"tags,omitempty" yaml:"tags" bson:"tags"`
	Priority    []KanbanPriorityType    `json:"priority,omitempty" yaml:"priority" bson:"priority"`
	Status      []KanbanStatusType      `json:"status,omitempty" yaml:"status" bson:"status"`
	DueDate     []KanbanDateDueType     `json:"due_date,omitempty" yaml:"due_date" bson:"due_date"`
	Dates       []KanbanDateType        `json:"dates,omitempty" yaml:"dates" bson:"dates"`
	Description []KanbanDescriptionType `json:"description,omitempty" yaml:"description" bson:"description"`
	Comments    []KanbanCommentType     `json:"comments,omitempty" yaml:"comments" bson:"comments"`
	Attachments []KanbanAttachmentType  `json:"attachments,omitempty" yaml:"attachments" bson:"attachments"`
	CheckList   []KanbanCheckListType   `json:"check_list,omitempty" yaml:"check_list" bson:"check_list"`
	Images      []KanbanImageType       `json:"images,omitempty" yaml:"images" bson:"images"`
	Archived    []KanbanArchivedType    `json:"archived,omitempty" yaml:"archived" bson:"archived"`
	Activity    []KanbanActivityType    `json:"activity,omitempty" yaml:"activity" bson:"activity"`
	Actions     []KanbanActionType      `json:"actions,omitempty" yaml:"actions" bson:"actions"`
}

type KanbanDescriptionType struct {
	Description string `json:"description,omitempty" yaml:"description" bson:"description"`
}

type KanbanPriorityType struct {
	UUID     string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Priority string `json:"priority" yaml:"priority" bson:"priority"`
}

type KanbanStatusType struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
}

type KanbanDateDueType struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Date string `json:"date" yaml:"date" bson:"date"`
}

type KanbanArchivedType struct {
	Archived bool   `json:"archived" yaml:"archived" bson:"archived"`
	Date     string `json:"date" yaml:"date" bson:"date"`
	User     string `json:"user" yaml:"user" bson:"user"`
}

type KanbanMemberType struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
}

type KanbanTagType struct {
	UUID  string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name  string `json:"name" yaml:"name" bson:"name"`
	Color string `json:"color" yaml:"color" bson:"color"`
	Icon  string `json:"icon,omitempty" yaml:"icon" bson:"icon"`
}

type KanbanDateType struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name string `json:"name" yaml:"name" bson:"name"`
	Date string `json:"date" yaml:"date" bson:"date"`
}

type KanbanAttachmentType struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name string `json:"name" yaml:"name" bson:"name"`
	Type string `json:"type" yaml:"type" bson:"type"`
	Size string `json:"size" yaml:"size" bson:"size"`
	Date string `json:"date" yaml:"date" bson:"date"`
	User string `json:"user" yaml:"user" bson:"user"`
}

type KanbanImageType struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name string `json:"name" yaml:"name" bson:"name"`
	Type string `json:"type" yaml:"type" bson:"type"`
	Size string `json:"size" yaml:"size" bson:"size"`
	Date string `json:"date" yaml:"date" bson:"date"`
	User string `json:"user" yaml:"user" bson:"user"`
	URL  string `json:"url" yaml:"url" bson:"url"`
}

type KanbanCommentType struct {
	UUID    string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Message string `json:"message" yaml:"message" bson:"message"`
	Date    string `json:"date" yaml:"date" bson:"date"`
	User    string `json:"user" yaml:"user" bson:"user"`
}

type KanbanCheckListType struct {
	UUID  string                    `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name  string                    `json:"name" yaml:"name" bson:"name"`
	Items []KanbanCheckListItemType `json:"items,omitempty" yaml:"items" bson:"items"`
}

type KanbanCheckListItemType struct {
	UUID    string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name    string `json:"name" yaml:"name" bson:"name"`
	Checked bool   `json:"checked" yaml:"checked" bson:"checked"`
}

type KanbanActionType struct {
	UUID   string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name   string `json:"name" yaml:"name" bson:"name"`
	Icon   string `json:"icon" yaml:"icon" bson:"icon"`
	Color  string `json:"color" yaml:"color" bson:"color"`
	Type   string `json:"type" yaml:"type" bson:"type"`
	Action string `json:"action" yaml:"action" bson:"action"`
}

type KanbanActivityType struct {
	UUID    string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Message string `json:"message" yaml:"message" bson:"message"`
	Date    string `json:"date" yaml:"date" bson:"date"`
	User    string `json:"user" yaml:"user" bson:"user"`
}
