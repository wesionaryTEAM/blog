package common

import (
	"fmt"
	"github/bhattaraibishal50/blog/config"
	"log"

	"github.com/jinzhu/gorm"
)

// DB database
var DB *gorm.DB

// DbConnection - Opening a database and save the reference to `Database` struct.
func DbConnection() *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DbUser, config.DbPassword, config.DbHost, config.Db)
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.DB().SetMaxIdleConns(10)
	DB = db
	return DB
}

// MigrateDatabase database tables
func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate()
}

// GetDatabase returns the db var
func GetDatabase() *gorm.DB {
	return DB
}
