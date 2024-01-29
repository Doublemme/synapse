package synapse

import (
	"time"

	"gorm.io/gorm"
	oauth2Gorm "src.techknowlogick.com/oauth2-gorm"

	"github.com/doublemme/synapse/pkg/core/middlewares"
	cm "github.com/doublemme/synapse/pkg/core/models"
	"github.com/doublemme/synapse/pkg/synapse/types"
	"github.com/labstack/echo/v4"
)

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
//
// @param ...modules - A list of functions of type LoadModuleFunc to load all the modules provided
func NewSynapseService(DbConnFunc types.DatabaseConnFunc, Options *SynapseOpts, modules ...types.LoadModuleFunc) *SynapseConfig {
	// Load default models
	defaultModels := append(make([]interface{}, 0),
		&cm.OauthUser{},
		&cm.AuthRole{},
		cm.AuthModule{},
		cm.AuthResource{},
		cm.AuthAction{},
	)

	conf := SynapseConfig{
		core: types.ModuleConfig{
			Acl:    make([]interface{}, 0),
			Models: defaultModels,
			Routes: make([]interface{}, 0),
		},
		DatabaseConn: DbConnFunc,
		Opts:         *Options,
	}

	// Load all the modules provided
	for _, m := range modules {
		conf.Modules = append(conf.Modules, m())
	}

	return &conf
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

	//Execute migrations
	err = db.AutoMigrate(sc.core.Models...)
	if err != nil {
		return err
	}

	defer tokenStore.Close()

	container := middlewares.ContainerContext{
		TokenConfig:  tokenStore,
		ClientConfig: clientStore,
	}

	e.Use(middlewares.ContainerMiddleware(&container))

	return nil
}

func (sc SynapseConfig) GetCoreConfig() types.ModuleConfig {
	return sc.core
}
