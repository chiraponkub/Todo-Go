package structdb

type GormDB interface {
	Register(data User) (Error error)
	DBGetUser(id string, isNull bool) (res []User, Error error)
	GetUsername(username string) (res User, Error error)
	UpdateUser(data User) (Error error)
	DelUser(id uint) (Error error)

	GetTodo(Id uint) (res []User, Error error)
	AddTodo(data Todo) (Error error)
	EditTodo(data Todo, role string) (Error error)
	DelTodo(id, Userid uint, role string) (Error error)
}
