package present

import (
	"ProjectEcho/access/constant"
	"ProjectEcho/control"
	"ProjectEcho/present/structure"
	"ProjectEcho/utility"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (Connect *CustomerHandler) getTodo(context echo.Context) error {
	api := control.APIControl{}
	Userid, role, err := GetId(context)
	paramId := context.Param("id")
	i, _ := strconv.Atoi(paramId)
	id := uint(i)
	if Userid != id && role == string(constant.Member) {
		return utility.ResponseError(context, http.StatusBadRequest, "ไอดีผู้ใช้งานไม่ถูกต้อง")
	}
	data, err := Connect.GetTodo(id)
	if err != nil {
		return utility.ResponseSuccess(context, http.StatusBadRequest, err.Error())
	}
	res, err := api.GetTodo(data)
	if err != nil {
		return utility.ResponseSuccess(context, http.StatusBadRequest, err.Error())
	}
	return context.JSON(http.StatusOK, res)
}

func (Connect *CustomerHandler) addTodo(context echo.Context) error {
	api := control.APIControl{}
	Userid, role, err := GetId(context)
	paramId := context.Param("id")
	i, _ := strconv.Atoi(paramId)
	id := uint(i)
	if Userid != id && role == string(constant.Member) {
		return utility.ResponseError(context, http.StatusBadRequest, "ไอดีผู้ใช้งานไม่ถูกต้อง")
	}
	req := structure.AddTodo{}
	err = context.Bind(&req)
	if err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	data, err := api.AddTodo(id, &req)
	if err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	err = Connect.AddTodo(data)
	if err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	return utility.ResponseSuccess(context, http.StatusOK, "Success")
}

func (Connect *CustomerHandler) editTodo(context echo.Context) error {
	api := control.APIControl{}
	Userid, role, _ := GetId(context)
	paramId := context.Param("id")
	i, _ := strconv.Atoi(paramId)
	TodoId := uint(i)
	req := structure.EditTodo{}
	if err := context.Bind(&req); err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	if err := ValidateStruct(&req); err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}

	data, err := api.EditTodo(TodoId, Userid, &req)
	if err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}

	if err = Connect.EditTodo(data, role); err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}

	return utility.ResponseSuccess(context, http.StatusOK, "Success")
}

func (Connect *CustomerHandler) delTodo(context echo.Context) error {
	Userid, role, _ := GetId(context)
	paramId := context.Param("id")
	i, _ := strconv.Atoi(paramId)
	id := uint(i)
	err := Connect.DelTodo(id, Userid, role)
	if err != nil {
		return utility.ResponseError(context, http.StatusBadRequest, err.Error())
	}
	return utility.ResponseSuccess(context, http.StatusOK, "Success")
}
