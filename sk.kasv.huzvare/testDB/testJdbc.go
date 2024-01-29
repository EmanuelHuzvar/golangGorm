package testDB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func IsJdbcRunning(dsn string) bool {
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
		return false
	} else {
		log.Println("database is running ")
		return true

	}

}
