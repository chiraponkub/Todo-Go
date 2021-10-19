package utility

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

func GetRoutes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
	}
	ioutil.WriteFile("routes.json", data, 0644)
}
