package connectors

import (
	"fmt"
	"imgverter/util"

	"github.com/konotorii/go-consola"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresInit() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		util.Config.DatabaseSettings.Host,
		util.Config.DatabaseSettings.User,
		util.Config.DatabaseSettings.Pass,
		util.Config.DatabaseSettings.Name,
		util.Config.DatabaseSettings.Port,
		util.Config.TimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		consola.Error("PG Database connection failed!", err)
		return nil
	}

	return db
}
