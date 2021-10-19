package structdb

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Password  string
	FirstName string
	LastName  string
	Role      string
	Todo []Todo `gorm:"foreignKey:UserRefer"`
}
