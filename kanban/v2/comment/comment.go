package comment

type Comment struct {
	User    string `json:"user"`
	UUID    string `json:"uuid"`
	Message string `json:"message"`
}
