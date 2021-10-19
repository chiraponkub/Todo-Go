package constant

import "errors"

const (
	LocalsKeyControl string = "CTRL"
	SecretKey string = "ToDoList_bell"
)

type UserRole string

const (
	Admin    UserRole = "admin"
	Member    UserRole = "Member"
)

var UserRoleData = []UserRole{
	Member,
	Admin,
}

func (userRole UserRole) Role() (result *string, Errors error) {
	switch userRole {
	case Member:
		fullName := "สมาชิก"
		result = &fullName
	case Admin:
		fullName := "ผู้ดูแลระบบ"
		result = &fullName
	default:
		Errors = errors.New("unimplemented")
	}
	return
}