package contracts

type AddToDoCommand struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	AssigneeId string `json:"assigneeId"`
}
