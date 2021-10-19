package utility

import "github.com/labstack/echo/v4"

func ResponseError(context echo.Context,statusCode int,message string) error {
	return context.JSON(statusCode,echo.Map{"message":message})
}

func ResponseSuccess(context echo.Context,statusCode int,message string) error {
	return context.JSON(statusCode,echo.Map{"message":message})
}