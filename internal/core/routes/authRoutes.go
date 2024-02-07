package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitAuthRoutes(e *echo.Echo) {
	grp := e.Group("/auth")
	grp.POST("/login", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
	// Request the reset password
	grp.POST("/reset-password", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
	// Update the user's password
	grp.PUT("/reset-password", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
	// Require Authenticated user
	grp.GET("/identity", func(c echo.Context) error { return c.String(http.StatusOK, "Hello, World!") })
}
