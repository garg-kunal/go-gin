package internal

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"fmt"
)

func InitDB() *gorm.DB{
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5440 sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!=nil{
		fmt.Print("Unable to connect to that database")
		return nil;
	}

	return db;
}