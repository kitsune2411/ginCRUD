package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=" + DB_HOST + " user=" + DB_USER + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=" + DB_PORT + " sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed Connect to Database")
	}

}
