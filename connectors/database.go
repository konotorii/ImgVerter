package connectors

import (
	"imgverter/types"
	"imgverter/util"

	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	if util.Config.DatabaseDriver == "SQLITE" {
		DB = SqliteInit()
	} else if util.Config.DatabaseDriver == "PSQL" {
		DB = PostgresInit()
	}

	DB.AutoMigrate(&types.DBImage{})
}
