package main

import (
	"github.com/doublemme/synapse/pkg/synapse"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConn() gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN: "Michael:123456@tcp(192.168.100.34:3306)/synapse_db?charset=utf8mb4&parseTime=True&loc=Local",
	})
}

func main() {
	e := echo.New()

	// Load all the configurations
	service := synapse.NewSynapseService(DbConn, &synapse.DefaultOptions)
	service.Init(e)

	e.Logger.Fatal(e.Start(":4040"))
}
