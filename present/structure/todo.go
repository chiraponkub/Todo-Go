package structure

type UserRefer struct {
	UserId uint   `json:"user_id"`
	Todo   []Todo `json:"todo"`
}

type Todo struct {
	Id       uint   `json:"id"`
	Text     string `json:"text"`
	IsActive bool   `json:"is_active"`
}

type AddTodo struct {
	Text   string `json:"text"`
}

type EditTodo struct {
	Text     string `json:"text"`
	IsActive bool   `json:"is_active"`
}
