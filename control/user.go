package control

import (
	"ProjectEcho/access/constant"
	"ProjectEcho/present/structdb"
	"ProjectEcho/present/structure"
	"ProjectEcho/utility"
	"errors"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (ctrl *APIControl) Login(login structure.Login, db structdb.User) (token string, Error error) {
	err := utility.VerifyPassword(db.Password, login.Password)
	if err != nil {
		Error = errors.New("รหัสผ่านไม่ถูกต้อง")
		return
	}
	resToken, err := utility.AuthenticationLogin(db.ID, db.Role)
	token = resToken
	return
}

func (ctrl *APIControl) GetUser(data []structdb.User) (res structure.GetUser, Error error) {
	for _, m1 := range data {
		res = structure.GetUser{
			Id:        m1.ID,
			Username:  m1.Username,
			FirstName: m1.FirstName,
			LastName:  m1.LastName,
			Role:      m1.Role,
		}
	}
	return
}

func (ctrl *APIControl) GetUserAll(data []structdb.User) (res []structure.GetUser, Error error) {
	for _, m1 := range data {
		resArray := structure.GetUser{
			Id:        m1.ID,
			Username:  m1.Username,
			FirstName: m1.FirstName,
			LastName:  m1.LastName,
			Role:      m1.Role,
		}
		res = append(res,resArray)
	}
	return
}

func (ctrl *APIControl) Register(reqRegister *structure.Register) (res structdb.User,Error error) {
	reqRegister.Username = strings.ToLower(reqRegister.Username)
	reqRegister.Username = strings.Trim(reqRegister.Username, "\t \n")
	reqRegister.Password = strings.Trim(reqRegister.Password, "\t \n")
	reqRegister.ConfirmPassword = strings.Trim(reqRegister.ConfirmPassword, "\t \n")
	reqRegister.FirstName = strings.Trim(reqRegister.FirstName, "\t \n")
	reqRegister.LastName = strings.Trim(reqRegister.LastName, "\t \n")
	user, err := regexp.MatchString("^[a-z0-9_-]{6,20}$", reqRegister.Username)
	if !user {
		Error = errors.New("username ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว และมีอักษรพิเศษได้แค่ _- เท่านั้น")
		return
	}
	pass, err := regexp.MatchString("^[a-zA-Z0-9_!-]{6,20}$", reqRegister.Password)
	if !pass {
		Error = errors.New("password ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว ต้องมีตัวพิมพ์ใหญ่และตัวพิมพ์เล็กตัวเลขและมีอักษรพิเศษได้แค่ !_- เท่านั้น")
		return
	}
	err = utility.ValidPassword(reqRegister.Password)
	if err != nil {
		Error = errors.New("password ต้องไม่ต่ำกว่า 6 ตัว และ ไม่เกิน 20 ตัว ต้องมีตัวพิมพ์ใหญ่และตัวพิมพ์เล็กตัวเลขและมีอักษรพิเศษได้แค่ !_- เท่านั้น")
		return
	}
	if reqRegister.Password != reqRegister.ConfirmPassword {
		Error = errors.New("รหัสผ่านไม่ตรงกัน")
		return
	}
	if !(len(reqRegister.FirstName) <= 30) {
		Error = errors.New("firstname ต้องไม่เกิน 30 ตัว")
		return
	}
	if !(len(reqRegister.LastName) <= 30) {
		Error = errors.New("lastname ต้องไม่เกิน 30 ตัว")
		return
	}
	if reqRegister.FirstName == "" {
		Error = errors.New("firstname ต้องไม่ว่าง")
		return
	}
	if reqRegister.LastName == "" {
		Error = errors.New("lastname ต้องไม่ว่าง")
		return
	}
	//err := ctrl.access.RDBMS.

	HashPassword, err := utility.Hash(reqRegister.Password)

	data := structdb.User{
		Username:  reqRegister.Username,
		Password:  string(HashPassword),
		FirstName: reqRegister.FirstName,
		LastName:  reqRegister.LastName,
		Role:      string(constant.Member),
	}
	res = data
	return
}

func (ctrl *APIControl) EditProfile(reqEditProfile *structure.EditUser,paramId string) (User structdb.User, Error error) {
	reqEditProfile.FirstName = strings.Trim(reqEditProfile.FirstName, "\t \n")
	reqEditProfile.LastName = strings.Trim(reqEditProfile.LastName, "\t \n")

	id ,_:= strconv.Atoi(paramId)

	res := structdb.User{
		Model:     gorm.Model{
			ID:        uint(id),
			UpdatedAt: time.Now(),
		},
		FirstName: reqEditProfile.FirstName,
		LastName:  reqEditProfile.LastName,
	}
	User = res
	return
}