package routes

import (
	"net/http"

	"github.com/doublemme/synapse/pkg/helpers"
	authView "github.com/doublemme/synapse/pkg/view/auth"
	"github.com/labstack/echo/v4"
)

func InitAuthRoutes(e *echo.Echo) {
	grp := e.Group("/auth")
	grp.GET("/login", func(c echo.Context) error { return helpers.Render(c, http.StatusOK, authView.Login()) })
	// Logout
	grp.POST("/logout", func(c echo.Context) error {
		c.Response().Header().Set("HX-Redirect", "http://localhost:3000/auth/login")
		return c.HTML(http.StatusOK, "<p style=\"color:red;\">Logged out</p>")
	})
	grp.POST("/login", func(c echo.Context) error {
		c.Response().Header().Set("HX-Redirect", "http://localhost:3000/dashboard")
		return c.NoContent(http.StatusOK)
	})
	// Request the reset password
	grp.POST("/reset-password", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
	// Update the user's password
	grp.PUT("/reset-password", func(c echo.Context) error { return c.JSON(http.StatusOK, new(interface{})) })
	// Require Authenticated user
	grp.GET("/identity", func(c echo.Context) error { return c.String(http.StatusOK, "Hello, World!") })
}
