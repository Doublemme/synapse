package synapse

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConn() gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN: "synapse:SynapseT3stDb!@tcp(localhost:3306)/synapse_db?charset=utf8mb4&parseTime=True&loc=Local",
	})
}

func TestSynapseService(t *testing.T) {

	e := echo.New()
	svc := NewSynapseService(DbConn, &DefaultOptions)

	assert.Nil(t, svc.Init(e), "error initializing the synapse service")

}
