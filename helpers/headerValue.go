package helpers

import (
	"github.com/labstack/echo/v4"
)

func GetContentType(ctx echo.Context) string{
	return ctx.Request().Header.Get("content-type")
}