package middlewares

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	oauth2Gorm "src.techknowlogick.com/oauth2-gorm"
)

type ContainerContext struct {
	echo.Context
	TokenConfig  *oauth2Gorm.TokenStore
	ClientConfig *oauth2Gorm.ClientStore
	Db           *gorm.DB
}

func ContainerMiddleware(container *ContainerContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			container.Context = c
			return next(container)
		}
	}
}
