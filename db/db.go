package db

import (
	"AvitoMessengeBot/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

const sqlite_db_name string = "clients.db"

func Init(db_type string) (db *gorm.DB) {
	var err error
	if db_type == "postgres" {
		db, err = gorm.Open(postgres.Open(config.Config.DSN), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}

	} else if db_type == "sqlite" {
		db, err = gorm.Open(sqlite.Open(sqlite_db_name), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatal("Error: wrong type of DB")
		return nil
	}
	err = db.AutoMigrate(&Client{}, &User{}, &Webhook{})
	if err != nil {
		log.Fatalln(err)
	}
	//db.Where("1=1").Delete(&models.Birthday{})
	//db.Where("1=1").Delete(&models.User{})
	return db
}
