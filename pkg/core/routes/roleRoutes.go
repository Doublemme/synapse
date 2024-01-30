package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRoleRoutes(e *echo.Echo) {
	grp := e.Group("/roles")
	// Retrieve a the list of roles
	grp.GET("/", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
	// Retrieve a role
	grp.GET("/:id", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
	// Create a new role
	grp.POST("/:id", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
	// Update the role's data
	grp.PUT("/:id", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
	// Delete a role
	grp.DELETE("/:id", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
}
