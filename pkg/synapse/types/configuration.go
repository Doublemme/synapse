package types

type ModuleConfig struct {
	Acl    []interface{}
	Models []interface{}
	Routes []interface{}
}

type LoadModuleFunc func() ModuleConfig

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
