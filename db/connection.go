package db

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var connection *gorm.DB

func GetConnection() *gorm.DB {
	return connection
}

func Init() {
	var err error

	connection, err = gorm.Open(sqlite.Open(viper.GetString("db.name")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
