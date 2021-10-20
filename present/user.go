package present

import (
	"github.com/chiraponkub/Todo-Go/access/constant"
	"github.com/chiraponkub/Todo-Go/control"
	"github.com/chiraponkub/Todo-Go/present/structdb"
	"github.com/chiraponkub/Todo-Go/present/structure"
	"github.com/chiraponkub/Todo-Go/utility"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetId(context echo.Context) (userId uint, role string, Error error) {
	User := context.Get("user").(*jwt.Token)
	claims := User.Claims.(jwt.MapClaims)
	var id = claims["id"].(float64)
	var userRole = claims["role"].(string)
	role = userRole
	userId = uint(id)
	return
}

func (Connect *CustomerHandler) admin(context echo.Context) error {

	hash, _ := utility.Hash("1234")
	admin := structdb.User{
		Username:  "admin",
		Password:  string(hash),
		FirstName: "",
		LastName:  "",
		Role:      string(constant.Admin),
	}
	_, err := Connect.GetUsername(admin.Username)
	if err == nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	if err = Connect.Register(admin); err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	return utility.ResponseSuccess(context, http.StatusOK, "สมัครสมาชิกสำเร็จ")
}

func (Connect *CustomerHandler) getUserAll(context echo.Context) error {
	api := control.APIControl{}
	data, err := Connect.DBGetUser("", false)
	if err != nil {
		return utility.ResponseError(context, http.StatusNotFound, err.Error())
	}
	res, err := api.GetUserAll(data)
	if err != nil {
		return utility.ResponseError(context, http.StatusNotFound, err.Error())
	}
	return context.JSON(http.StatusOK, res)
}

func (Connect *CustomerHandler) editProfile(context echo.Context) error {
	api := control.APIControl{}
	Userid, role, err := GetId(context)
	paramId := context.Param("id")
	i, _ := strconv.Atoi(paramId)
	id := uint(i)
	if Userid != id && role == string(constant.Member) {
		return utility.ResponseError(context, http.StatusBadRequest, "ไอดีผู้ใช้งานไม่ถูกต้อง")
	}
	reqEditProfile := new(structure.EditUser)
	if err = context.Bind(reqEditProfile); err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	err = ValidateStruct(reqEditProfile)
	if err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	data, err := api.EditProfile(reqEditProfile, paramId)
	if err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	err = Connect.UpdateUser(data)
	if err != nil {
		return utility.ResponseError(context, http.StatusInternalServerError, err.Error())
	}
	return utility.ResponseSuccess(context, http.StatusOK, "แก้ไขสำเร็จ")
}

func (Connect *CustomerHandler) delUser(context echo.Context) error {
	id := context.Param("id")
	UserId, _ := strconv.Atoi(id)
	err := Connect.DelUser(uint(UserId))
	if err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}

	return utility.ResponseSuccess(context, http.StatusOK, "Success")
}

func (Connect *CustomerHandler) GetUser(context echo.Context) error {
	api := control.APIControl{}
	Userid, role, err := GetId(context)
	paramId := context.Param("id")
	i, _ := strconv.Atoi(paramId)
	id := uint(i)
	if Userid != id && role == string(constant.Member) {
		return utility.ResponseError(context, http.StatusBadRequest, "ไอดีผู้ใช้งานไม่ถูกต้อง")
	}
	var res structure.GetUser
	data, err := Connect.DBGetUser(paramId, true)
	if err != nil {
		return utility.ResponseError(context, http.StatusNotFound, err.Error())
	}
	res, err = api.GetUser(data)
	if err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	return context.JSON(http.StatusOK, res)
}

func (Connect *CustomerHandler) register(context echo.Context) error {
	api := control.APIControl{}
	data := structdb.User{}
	reqRegister := new(structure.Register)
	if err := context.Bind(&reqRegister); err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	err := ValidateStruct(&reqRegister)
	if err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	_, err = Connect.GetUsername(data.Username)
	if err == nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	User, err := api.Register(reqRegister)
	if err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	if err = Connect.Register(User); err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	return utility.ResponseSuccess(context, http.StatusOK, "สมัครสมาชิกสำเร็จ")
}

func (Connect *CustomerHandler) login(context echo.Context) error {
	api := control.APIControl{}
	login := structure.Login{}
	if err := context.Bind(&login); err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	err := ValidateStruct(&login)
	if err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	data, err := Connect.GetUsername(login.Username)
	token, err := api.Login(login, data)
	if err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	return utility.ResponseSuccess(context, http.StatusOK, token)
}
