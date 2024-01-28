package database

import (
	"bookapp/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database configurations
// Change based on your MySQL configurations
const DB_USERNAME = "root"
const DB_PASSWORD = "root"
const DB_NAME = "bookapp"
const DB_HOST = "localhost"
const DB_PORT = "3306"

var DB *gorm.DB

func InitDb() (*gorm.DB, error) {
	DB = connectDB()

	//Create the Book Table
	err := DB.AutoMigrate(&models.Book{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}

// Establish the connection with MySQL
func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
