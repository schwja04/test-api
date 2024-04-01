package contracts

type UpdateToDoCommand struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	AssigneeId string `json:"assigneeId"`
}
