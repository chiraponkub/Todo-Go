package structdb

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Text      string
	IsActive  bool
	UserRefer uint
}


