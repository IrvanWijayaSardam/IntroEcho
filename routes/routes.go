package routes

import (
	"restEcho1/controller"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uc controller.UserController) {
	e.POST("/users", uc.Create())
	e.GET("/users", uc.GetUsers())
	e.GET("/users/:id", uc.Get())
	e.DELETE("/users", uc.Delete())
	e.PUT("/users/:id", uc.Update())
}
