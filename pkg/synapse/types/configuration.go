package types

import "github.com/labstack/echo/v4"

type ModuleConfig struct {
	Acl    []AclModule
	Models []interface{}
	Routes []InitModuleRoutes
}

type LoadModuleFunc func() ModuleConfig
type InitModuleRoutes func(e *echo.Echo)

type AclModule struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Resources   []AclResource `json:"resources"`
}

type AclResource struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Actions     []AclAction `json:"actions"`
}

type AclAction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
