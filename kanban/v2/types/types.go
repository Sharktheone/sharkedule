package types

type Task struct {
	Name         string    `json:"name"`
	UUID         string    `json:"uuid"`
	Boards       []string  `json:"boards"`
	Columns      []string  `json:"columns"`
	Tags         []string  `json:"tags"`
	Dependencies []string  `json:"dependencies"`
	Dependents   []string  `json:"dependents"`
	Comments     []Comment `json:"comments"`
	Description  string    `json:"description"`
	Members      []string  `json:"members"`
	Priority     string    `json:"priority"`
	Status       string    `json:"status"`
	DueDate      string    `json:"due_date"`
	Dates        []string  `json:"dates"`
	Attachments  []string  `json:"attachments"`
	CheckList    []string  `json:"check_list"`
	Done         bool      `json:"done"`
	Activity     []string  `json:"activity"`
}

type Column struct {
	Name        string   `json:"name"`
	UUID        string   `json:"uuid"`
	Boards      []string `json:"boards"`
	Tasks       []string `json:"tasks"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
}

type Tag struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	Color       string `json:"color"`
	Icon        string `json:"icon"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Board struct {
	Name        string     `json:"name"`
	UUID        string     `json:"uuid"`
	Columns     []string   `json:"columns"`
	Tags        []string   `json:"tags"`
	Description string     `json:"description"`
	Members     []string   `json:"members"`
	Priority    string     `json:"priority"`
	Status      string     `json:"status"`
	DueDate     string     `json:"due_date"`
	Dates       []string   `json:"dates"`
	Comments    []Comment  `json:"comments"`
	Attachments []string   `json:"attachments"`
	CheckList   []string   `json:"check_list"`
	Archived    bool       `json:"archived"`
	Activity    []Activity `json:"activity"`
	Actions     []string   `json:"actions"`
}

type Member struct {
	Username       string `json:"username"`
	UUID           string `json:"uuid"`
	ProfilePicture string `json:"profile_picture"`
}

type Checklist struct {
	Name        string          `json:"name"`
	UUID        string          `json:"uuid"`
	Items       []ChecklistItem `json:"items"`
	Description string          `json:"description"`
}

type ChecklistItem struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	Checked bool   `json:"checked"`
}

type Comment struct {
	User    string `json:"from"`
	UUID    string `json:"uuid"`
	Message string `json:"message"`
}

type Priority struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

type Status struct {
	Name        string `json:"name"`
	UUID        string `json:"uuid"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

type Attachment struct {
	UUID        string `json:"uuid"`
	User        string `json:"from"`
	Size        int64  `json:"size"`
	Type        string `json:"type"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type Date struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Date        int64  `json:"date"`
	Description string `json:"description"`
}

type Activity struct {
	//TODO
}

type BoardEnvironment struct {
	Tags        []Tag        `json:"tags"`
	Status      []Status     `json:"status"`
	Priority    []Priority   `json:"priority"`
	Columns     []Column     `json:"columns"`
	Boards      []Board      `json:"boards"`
	Tasks       []Task       `json:"tasks"`
	Members     []Member     `json:"members"`
	Checklist   []Checklist  `json:"checklist"`
	Attachments []Attachment `json:"attachments"`
	Dates       []Date       `json:"dates"`
}
