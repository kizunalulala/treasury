package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

var self *database

func Init() {
	var success bool
	for i := 0; i < 10; i++ {
		//dsn := fmt.Sprintf("root:password@tcp(127.0.0.1:3306)/treasury_db?charset=utf8mb4&parseTime=True&loc=Local")
		dsn := fmt.Sprintf("root:password@tcp(treasury_db:3306)/treasury_db?charset=utf8mb4&parseTime=True&loc=Local")
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		self = &database{db: db}
		if err == nil {
			success = true
			break
		} else {
			time.Sleep(10 * time.Second)
		}
	}
	if !success {
		log.Fatalf("Database connection failed")
	}
}
