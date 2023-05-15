package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"shopnilashik/entities"
)

var Database *gorm.DB

func Connect() error {
	var err error
	Database, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=dhaka1214 dbname=go_blog port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	fmt.Println("db.......", Database)
	if err != nil {
		panic(err)
	}
	fmt.Println("Established a successful connection!")
	Database.AutoMigrate(&entities.Post{})
	return nil
}
