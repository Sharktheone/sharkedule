package types

type Task struct {
	Name         string   `json:"name"`
	UUID         string   `json:"uuid"`
	Boards       []string `json:"boards"`
	Tags         []string `json:"tags"`
	Dependencies []string `json:"dependencies"`
	Dependents   []string `json:"dependents"`
}

type Column struct {
	Name  string   `json:"name"`
	UUID  string   `json:"uuid"`
	Board string   `json:"board"`
	Tasks []string `json:"tasks"`
	Tags  []string `json:"tags"`
}

type Tag struct {
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
	Color string `json:"color"`
	Icon  string `json:"icon"`
	Type  string `json:"type"`
}

type Board struct {
	Name    string   `json:"name"`
	UUID    string   `json:"uuid"`
	Columns []string `json:"columns"`
	Tags    []string `json:"tags"`
}
