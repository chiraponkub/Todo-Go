package present

import (
	"errors"
	"fmt"
	"github.com/chiraponkub/Todo-Go/access/constant"
	"github.com/chiraponkub/Todo-Go/control"
	"github.com/chiraponkub/Todo-Go/environment"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

type ContextApi struct {
	env        *environment.Properties
	apiControl *control.APIControl
}

func APICreate(*control.APIControl) {
	e := echo.New()

	prop := environment.Build()
	if prop == nil {
		log.Panic("environment not exist")
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	h := CustomerHandler{}
	h.Initialize(prop)

	api := e.Group("/api")
	api.POST("/register", h.register)
	api.POST("/login", h.login)

	user := api.Group("/user")
	user.GET("/getAll", h.getUserAll)
	user.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SuccessHandler: func(context echo.Context) {
			User := context.Get("user").(*jwt.Token)
			claims := User.Claims.(jwt.MapClaims)
			var userRole = claims["role"]
			if userRole == string(constant.Member) {

			} else {
				log.Panic(echo.NewHTTPError(http.StatusUnauthorized, "ไม่มีสิทธิ์ในการเข้าถึง"))
			}
		},
		ErrorHandler: AuthError,
		SigningKey:   []byte(constant.SecretKey),
		AuthScheme:   "Bearer",
	}))

	user.GET("/:id", h.GetUser)
	user.PUT("/:id", h.editProfile)

	todo := user.Group("/todo")
	todo.GET("/:id", h.getTodo)
	todo.POST("/:id", h.addTodo)
	todo.PUT("/:id", h.editTodo)
	todo.DELETE("/:id", h.delTodo)

	admin := api.Group("/admin")
	admin.POST("/register", h.admin)
	admin.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(constant.SecretKey),
		SuccessHandler: func(context echo.Context) {
			User := context.Get("user").(*jwt.Token)
			claims := User.Claims.(jwt.MapClaims)
			var userRole = claims["role"].(string)
			if userRole == string(constant.Admin) {

			} else {
				log.Panic(echo.NewHTTPError(http.StatusUnauthorized, "ไม่มีสิทธิ์ในการเข้าถึง"))
			}
		},
		ErrorHandler: AuthError,
		AuthScheme:   "Bearer",
	}))

	admin.GET("/todo/:id", h.getTodo)
	admin.POST("/todo/:id", h.addTodo)
	admin.PUT("/todo/:id", h.editTodo)
	admin.DELETE("/todo/:id", h.delTodo)

	admin.GET("/user", h.getUserAll)
	admin.GET("/user/:id", h.GetUser)
	admin.PUT("/user/:id", h.editProfile)
	admin.DELETE("/user/:id", h.delUser)

	e.Logger.Fatal(e.Start(":1324"))
}

func ValidateStruct(dataStruct interface{}) error {
	validate := validator.New()
	err := validate.Struct(dataStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(fmt.Sprintf("%s: %s", err.StructField(), err.Tag()))
		}
	} else {
		return nil
	}
	return err
}

func AuthError(err error) (Error error) {
	Error = echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	return
}
