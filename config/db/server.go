package db

import (
	"os"

	"lms/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Server() *gorm.DB {
	// Load .env file
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	// Get the database credentials from .env file
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	/// Create the database Serverion string
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"

	// Server to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// make logger
		// Logger: logger.Default.LogMode(logger.Info),

		//logger off
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
	})

	// Auto migrate the database
	db.AutoMigrate(&model.User{}, &model.Knowledge{},
		&model.Course{}, &model.Reference{}, &model.Content{}, model.Section{},
		model.Category{}, model.Menu{}, model.Role{}, model.Quiz{}, model.Answer{},
		model.Question{})

	if err != nil {
		panic("Failed to Server to database!")
	}

	if err != nil {
		panic("Failed to Server to database!")
	} else {
		println("Servered to database!")
	}

	return db

}
