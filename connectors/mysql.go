package connectors

import (
	"fmt"
	"imgverter/util"

	"github.com/konotorii/go-consola"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySqlInit() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		util.Config.DatabaseSettings.User,
		util.Config.DatabaseSettings.Pass,
		util.Config.DatabaseSettings.Host,
		util.Config.DatabaseSettings.Port,
		util.Config.DatabaseSettings.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		consola.Error("MySQL Database connection failed!", err)
		return nil
	}

	return db
}
