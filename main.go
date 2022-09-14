package main

import (
	"fmt"
	"project/kutsuya/config"
	"project/kutsuya/factory"
	"project/kutsuya/migration"
	"project/kutsuya/utils/database/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitMysqlDB(cfg)
	e := echo.New()
	e.Use(middleware.CORS())
	migration.InitMigrate(db)
	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))

}
