package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo) {
	grp := e.Group("/users")
	// Retrieve a the list of users
	grp.GET("/", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
	// Retrieve a user
	grp.GET("/:id", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
	// Create a new user
	grp.POST("/:id", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
	// Update the user's data
	grp.PUT("/:id", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
	// Delete a user
	grp.DELETE("/:id", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
}
