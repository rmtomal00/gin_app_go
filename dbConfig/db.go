package db

import (
	"fmt"
	"os"

	//userModel "app.team71.link/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDb() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	});

	if err !=nil {
		fmt.Print(err.Error())
	}
	// var userModel userModel.User;
	// err = db.AutoMigrate(userModel);

	// if err != nil {
	// 	fmt.Print(err)
	// }
	//fmt.Println("Connect Db")
	return db
}