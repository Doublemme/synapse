package synapse

import (
	"os"
	"time"

	"gorm.io/gorm"
	oauth2Gorm "src.techknowlogick.com/oauth2-gorm"

	"github.com/doublemme/synapse/internal/core/middlewares"
	cm "github.com/doublemme/synapse/internal/core/models"
	"github.com/doublemme/synapse/internal/core/routes"
	"github.com/doublemme/synapse/internal/synapse/helpers"
	"github.com/doublemme/synapse/internal/synapse/types"
	"github.com/labstack/echo/v4"
)

type LoadModuleFunc func() types.ModuleConfig

type SynapseOpts struct {
	TokenLifetime       time.Duration
	ClientTokenLifetime time.Duration
}

type SynapseConfig struct {
	// The core module configuration
	core types.ModuleConfig
	// List of modules injected and that need to be load at startup
	Modules []types.ModuleConfig
	// Function that handle the database connection and return the Gorm dialector
	DatabaseConn types.DatabaseConnFunc
	Opts         SynapseOpts
}

var DefaultOptions SynapseOpts = SynapseOpts{
	TokenLifetime:       1 * time.Hour,
	ClientTokenLifetime: 1 * time.Hour,
}

// # Create a new Synapse service.
//
// @param DbConnFunc - A function that must return a `gorm.Dialector` that handle the connection to the database.
//
// @param Options    - Handle the options of the service
func NewSynapseService(DbConnFunc types.DatabaseConnFunc, Options *SynapseOpts) (*SynapseConfig, error) {
	// Load default models
	defaultModels := append(make([]interface{}, 0),
		&cm.OauthUser{},
		&cm.AuthRole{},
		cm.AuthModule{},
		cm.AuthResource{},
		cm.AuthAction{},
	)

	var coreAcl []types.AclModule

	file, err := os.OpenFile("../core/config/acl.json", 0, os.ModeAppend)
	if err != nil {
		return nil, err
	}

	err = helpers.UnmarshalAcl(&coreAcl, file)
	if err != nil {
		return nil, err
	}
	file.Close()

	coreConfig := types.ModuleConfig{
		Acl:    coreAcl,
		Models: defaultModels,
		Routes: []types.InitModuleRoutes{routes.InitAuthRoutes, routes.InitUserRoutes, routes.InitRoleRoutes},
	}

	conf := SynapseConfig{
		core:         coreConfig,
		DatabaseConn: DbConnFunc,
		Opts:         *Options,
	}

	return &conf, nil
}

// Initialize all the components in the configuration
// like the Database service, run the migrations, load all the middlewares from each module,
// load all the routes from every module etc. etc
func (sc *SynapseConfig) Init(e *echo.Echo) error {

	dialector := sc.DatabaseConn()

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}

	tokenStoreCfg := &oauth2Gorm.Config{
		TableName:   "oauth_token",
		MaxLifetime: sc.Opts.TokenLifetime,
		Dialector:   dialector,
	}
	clientStoreCfg := &oauth2Gorm.Config{
		TableName:   "oauth_client",
		MaxLifetime: sc.Opts.ClientTokenLifetime,
		Dialector:   dialector,
	}

	tokenStore := oauth2Gorm.NewTokenStoreWithDB(tokenStoreCfg, db, 0)
	clientStore := oauth2Gorm.NewClientStoreWithDB(clientStoreCfg, db)

	//Execute core migrations
	err = db.AutoMigrate(sc.core.Models...)
	if err != nil {
		return err
	}

	//Init the core routes
	for _, routeFunc := range sc.core.Routes {
		routeFunc(e)
	}

	//Load the core acl
	if err := helpers.SyncAcl(db, &sc.core.Acl); err != nil {
		return err
	}

	for _, module := range sc.Modules {

		//Execute modules migrations
		err = db.AutoMigrate(module.Models...)
		if err != nil {
			return err
		}

		//Init routes
		for _, routeFunc := range module.Routes {
			routeFunc(e)
		}

		if err := helpers.SyncAcl(db, &module.Acl); err != nil {
			return err
		}

	}

	defer tokenStore.Close()

	container := middlewares.ContainerContext{
		TokenConfig:  tokenStore,
		ClientConfig: clientStore,
		Db:           db,
	}

	e.Use(middlewares.ContainerMiddleware(&container))

	return nil
}

// Load the given modules into the Synapse service config.
//
// This must be invoked before the synapseConfig.Init()
func (sc *SynapseConfig) LoadModules(modules ...LoadModuleFunc) error {

	// Load all the modules provided
	for _, moduleFunc := range modules {
		sc.Modules = append(sc.Modules, moduleFunc())
	}

	return nil
}

func (sc SynapseConfig) GetCoreConfig() types.ModuleConfig {
	return sc.core
}
