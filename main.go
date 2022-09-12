package main

import (
	"fmt"
	"project/kutsuya/config"
	"project/kutsuya/factory"
	"project/kutsuya/migration"
	"project/kutsuya/utils/database/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitMysqlDB(cfg)
	e := echo.New()
	migration.InitMigrate(db)
	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))

}
