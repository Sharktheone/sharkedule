package kanbanboard

type KanbanBoard struct {
	UUID        string                 `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string                 `json:"name" yaml:"name" bson:"name"`
	Description kanbanDescriptionType  `json:"description,omitempty" yaml:"description" bson:"description"`
	Members     []kanbanMemberType     `json:"members,omitempty" yaml:"members" bson:"members"`
	Tags        []kanbanTagType        `json:"tags,omitempty" yaml:"tags" bson:"tags"`
	Priority    kanbanPriorityType     `json:"priority,omitempty" yaml:"priority" bson:"priority"`
	Status      kanbanStatusType       `json:"status,omitempty" yaml:"status" bson:"status"`
	DueDate     kanbanDateDueType      `json:"due_date,omitempty" yaml:"due_date" bson:"due_date"`
	Dates       []kanbanDateType       `json:"dates,omitempty" yaml:"dates" bson:"dates"`
	Comments    []kanbanCommentType    `json:"comments,omitempty" yaml:"comments" bson:"comments"`
	Attachments []kanbanAttachmentType `json:"attachments,omitempty" yaml:"attachments" bson:"attachments"`
	CheckList   []kanbanCheckListType  `json:"check_list,omitempty" yaml:"check_list" bson:"check_list"`
	Images      []kanbanImageType      `json:"images,omitempty" yaml:"images" bson:"images"`
	Archived    kanbanArchivedType     `json:"archived,omitempty" yaml:"archived" bson:"archived"`
	Activity    []kanbanActivityType   `json:"activity,omitempty" yaml:"activity" bson:"activity"`
	Actions     []kanbanActionType     `json:"actions,omitempty" yaml:"actions" bson:"actions"`
	Columns     []kanbanColumnType     `json:"columns,omitempty" yaml:"columns" bson:"columns"`
}

type kanbanColumnType struct {
	UUID        string                `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string                `json:"name" yaml:"name" bson:"name"`
	Description kanbanDescriptionType `json:"description,omitempty" yaml:"description" bson:"description"`
	Tasks       []kanbanTaskType      `json:"tasks,omitempty" yaml:"tasks" bson:"tasks"`
}

type kanbanTaskType struct {
	UUID        string                  `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name        string                  `json:"name" yaml:"name" bson:"name"`
	Type        string                  `json:"type" yaml:"type" bson:"type"`
	Members     []kanbanMemberType      `json:"members,omitempty" yaml:"members" bson:"members"`
	Tags        []kanbanTagType         `json:"tags,omitempty" yaml:"tags" bson:"tags"`
	Priority    []kanbanPriorityType    `json:"priority,omitempty" yaml:"priority" bson:"priority"`
	Status      []kanbanStatusType      `json:"status,omitempty" yaml:"status" bson:"status"`
	DueDate     []kanbanDateDueType     `json:"due_date,omitempty" yaml:"due_date" bson:"due_date"`
	Dates       []kanbanDateType        `json:"dates,omitempty" yaml:"dates" bson:"dates"`
	Description []kanbanDescriptionType `json:"description,omitempty" yaml:"description" bson:"description"`
	Comments    []kanbanCommentType     `json:"comments,omitempty" yaml:"comments" bson:"comments"`
	Attachments []kanbanAttachmentType  `json:"attachments,omitempty" yaml:"attachments" bson:"attachments"`
	CheckList   []kanbanCheckListType   `json:"check_list,omitempty" yaml:"check_list" bson:"check_list"`
	Images      []kanbanImageType       `json:"images,omitempty" yaml:"images" bson:"images"`
	Archived    []kanbanArchivedType    `json:"archived,omitempty" yaml:"archived" bson:"archived"`
	Activity    []kanbanActivityType    `json:"activity,omitempty" yaml:"activity" bson:"activity"`
	Actions     []kanbanActionType      `json:"actions,omitempty" yaml:"actions" bson:"actions"`
}

type kanbanDescriptionType struct {
	Description string `json:"description,omitempty" yaml:"description" bson:"description"`
}

type kanbanPriorityType struct {
	UUID     string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Priority string `json:"priority" yaml:"priority" bson:"priority"`
}

type kanbanStatusType struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
}

type kanbanDateDueType struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Date string `json:"date" yaml:"date" bson:"date"`
}

type kanbanArchivedType struct {
	Archived bool   `json:"archived" yaml:"archived" bson:"archived"`
	Date     string `json:"date" yaml:"date" bson:"date"`
	User     string `json:"user" yaml:"user" bson:"user"`
}

type kanbanMemberType struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
}

type kanbanTagType struct {
	UUID  string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name  string `json:"name" yaml:"name" bson:"name"`
	Color string `json:"color" yaml:"color" bson:"color"`
	Icon  string `json:"icon,omitempty" yaml:"icon" bson:"icon"`
}

type kanbanDateType struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name string `json:"name" yaml:"name" bson:"name"`
	Date string `json:"date" yaml:"date" bson:"date"`
}

type kanbanAttachmentType struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name string `json:"name" yaml:"name" bson:"name"`
	Type string `json:"type" yaml:"type" bson:"type"`
	Size string `json:"size" yaml:"size" bson:"size"`
	Date string `json:"date" yaml:"date" bson:"date"`
	User string `json:"user" yaml:"user" bson:"user"`
}

type kanbanImageType struct {
	UUID string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name string `json:"name" yaml:"name" bson:"name"`
	Type string `json:"type" yaml:"type" bson:"type"`
	Size string `json:"size" yaml:"size" bson:"size"`
	Date string `json:"date" yaml:"date" bson:"date"`
	User string `json:"user" yaml:"user" bson:"user"`
	URL  string `json:"url" yaml:"url" bson:"url"`
}

type kanbanCommentType struct {
	UUID    string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Message string `json:"message" yaml:"message" bson:"message"`
	Date    string `json:"date" yaml:"date" bson:"date"`
	User    string `json:"user" yaml:"user" bson:"user"`
}

type kanbanCheckListType struct {
	UUID  string                    `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name  string                    `json:"name" yaml:"name" bson:"name"`
	Items []kanbanCheckListItemType `json:"items,omitempty" yaml:"items" bson:"items"`
}

type kanbanCheckListItemType struct {
	UUID    string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name    string `json:"name" yaml:"name" bson:"name"`
	Checked bool   `json:"checked" yaml:"checked" bson:"checked"`
}

type kanbanActionType struct {
	UUID   string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Name   string `json:"name" yaml:"name" bson:"name"`
	Icon   string `json:"icon" yaml:"icon" bson:"icon"`
	Color  string `json:"color" yaml:"color" bson:"color"`
	Type   string `json:"type" yaml:"type" bson:"type"`
	Action string `json:"action" yaml:"action" bson:"action"`
}

type kanbanActivityType struct {
	UUID    string `json:"uuid" yaml:"uuid" bson:"uuid"`
	Message string `json:"message" yaml:"message" bson:"message"`
	Date    string `json:"date" yaml:"date" bson:"date"`
	User    string `json:"user" yaml:"user" bson:"user"`
}
