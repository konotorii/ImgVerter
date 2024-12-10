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
	}

	DB.AutoMigrate(&types.DBImage{})
}
