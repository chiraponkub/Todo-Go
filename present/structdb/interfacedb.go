package structdb

type GormDB interface {
	Register(data User) (Error error)
	GetUser(id string,isNull bool) (res []User, Error error)
	GetUsername(username string) (res User,Error error)
	UpdateUser(data User) (Error error)
	DelUser(id uint) (Error error)

	GetTodo(Id uint) (res []Todo,Error error)
	AddTodo(data Todo) (Error error)
	EditTodo(data Todo, role string) (Error error)
	DelTodo(id ,Userid uint, role string) (Error error)
}
