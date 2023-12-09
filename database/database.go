package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/nchdatta/go-note-app/config"
	"github.com/nchdatta/go-note-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

// Function to connect to the database
func ConnectDB() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 16)

	if err != nil {
		log.Fatalln("Failed to get database port!")
	}

	// Createing Data Source Name (DSN) dynamically
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to the database.\n", err)
	}

	log.Println("DB connected!")
	errMig := db.AutoMigrate(&models.Note{})
	if errMig != nil {
		log.Fatalln("Error Occured on migrating.\n", errMig)
	}
	log.Println("DB migrated!")
	DBConn = db

}
