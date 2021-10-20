package present

import (
	"errors"
	"fmt"
	"github.com/chiraponkub/Todo-Go/access/constant"
	"github.com/chiraponkub/Todo-Go/environment"
	"github.com/chiraponkub/Todo-Go/present/structdb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerHandler struct {
	DB *gorm.DB
}

func (Connect *CustomerHandler) Initialize(env *environment.Properties) {
	databaseSet := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		env.GormHost, env.GormPort, env.GormUserDB, env.GormNameDB, env.GormPassDB, "disable")

	db, err := gorm.Open(postgres.Open(databaseSet), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database : %s", err.Error()))
	}

	if env.Flavor != environment.Production {
		db = db.Debug()
	}

	_ = db.AutoMigrate(
		&structdb.User{},
	)
	Connect.DB = db
}

// user

func (Connect CustomerHandler) Register(data structdb.User) (Error error) {
	err := Connect.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}

func (Connect CustomerHandler) GetUsername(username string) (res structdb.User, Error error) {
	var data structdb.User
	db := Connect.DB
	err := db.Where("username = ?", username).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Error = errors.New("ไม่มีผู้ใช้คนนี้อยู่ในระบบ")
			return
		}
	}
	res = data
	return
}

func (Connect CustomerHandler) DBGetUser(id string, isNull bool) (res []structdb.User, Error error) {
	var data []structdb.User
	db := Connect.DB
	if isNull == true {
		err := db.Where("id = ? and role = ?", id, string(constant.Member)).First(&data).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				Error = err
				return
			}
		}
	} else {
		if err := db.Where("role = ?", string(constant.Member)).Find(&data).Error; err != nil {
			Error = err
			return
		}
	}
	res = data
	return
}

func (Connect CustomerHandler) UpdateUser(data structdb.User) (Error error) {
	err := Connect.DB.Where("id = ?", data.ID).Updates(&data).Error
	if err != nil {
		Error = err
		return
	}
	return
}

func (Connect CustomerHandler) DelUser(id uint) (Error error) {
	db := Connect.DB
	data := structdb.User{}
	if err := db.Where("id = ? ", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Error = err
			return
		}
	}
	err := db.Select("Todo").Delete(&data).Error
	if err != nil {
		Error = err
		return
	}
	return
}

// todo

func (Connect CustomerHandler) GetTodo(Userid uint) (res []structdb.User, Error error) {
	var Data []structdb.User
	db := Connect.DB
	err := db.Where("id = ? and role != ? ", Userid, string(constant.Admin)).Preload("Todo").Find(&Data).Error
	if err != nil {
		Error = err
		return
	}
	res = Data
	return
}

func (Connect CustomerHandler) AddTodo(data structdb.Todo) (Error error) {
	err := Connect.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&data).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRegistered) {
			Error = err
			return
		}
	}
	return
}

func (Connect CustomerHandler) EditTodo(data structdb.Todo, role string) (Error error) {
	db := Connect.DB
	res := structdb.Todo{}
	if role == string(constant.Member) {
		err := db.Where("id = ? And user_refer = ?", data.ID, data.UserRefer).Take(&res).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				Error = errors.New("ไม่พบสิ่งที่จะทำการแก้ไข")
				return
			}
		}
		res.Text = data.Text
		res.IsActive = data.IsActive
		err = db.Save(&res).Error
		if err != nil {
			Error = err
			return
		}
	}

	if role == string(constant.Admin) {
		err := db.Where("id = ?", data.ID).Take(&res).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				Error = errors.New("ไม่พบสิ่งที่จะทำการแก้ไข")
				return
			}
		}
		res.Text = data.Text
		res.IsActive = data.IsActive
		err = db.Save(&res).Error
		if err != nil {
			Error = err
			return
		}
	}
	return
}

func (Connect CustomerHandler) DelTodo(id, Userid uint, role string) (Error error) {
	db := Connect.DB
	data := structdb.Todo{}
	if role == string(constant.Member) {
		err := db.Where("id = ? And user_refer = ?", id, Userid).First(&data).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				Error = errors.New("ไม่พบสิ่งที่จะทำการลบ")
				return
			}
		}
		err = db.Where("id = ? And user_refer = ?", id, Userid).Delete(&data).Error
		if err != nil {
			Error = err
			return
		}
	}

	if role == string(constant.Admin) {
		err := db.Where("id = ?", id).Delete(&data).Error
		if err != nil {
			Error = err
			return
		}
	}
	return
}
