package connectors

import (
	"imgverter/util"

	"github.com/konotorii/go-consola"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SqliteInit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(util.Config.DatabaseSettings.Path), &gorm.Config{})
	if err != nil {
		consola.Error("SQLite database couldn't be opened!")
		return nil
	}

	return db
}
