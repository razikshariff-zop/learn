package config

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Api=os.Getenv("API_KEY")
func ConnectDB() {
	// Load from env vars
	username := ""
	rawPassword := ""
	password := url.QueryEscape(rawPassword)
	host := "sql12.freesqldatabase.com"
	port := "3306"
	dbname := "sql12794681"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	// Open DB connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ failed to connect to database: %v", err)
	}

	DB = db
	log.Println("✅ Database connected successfully")
}
