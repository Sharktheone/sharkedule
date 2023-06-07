package column

type Column struct {
	Name        string   `json:"name"`
	UUID        string   `json:"uuid"`
	Boards      []string `json:"boards"`
	Tasks       []string `json:"tasks"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
}
