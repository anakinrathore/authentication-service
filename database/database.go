package database

import (
    "fmt"
    // "go-authentication-boilerplate/models"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

// This will give us access to db connection when we import the database package
var DB *gorm.DB

// ConnectToDB connects the server with database
func ConnectToDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading env file \n", err)
    }

    dsn := fmt.Sprintf("host=localhost dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
                             os.Getenv("PSQL_DBNAME"), os.Getenv("PSQL_PORT"))

    log.Print("Connecting to PostgreSQL DB...")
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

    if err != nil {
        log.Fatal("Failed to connect to database. \n", err)
        os.Exit(2)

    }
    log.Println("connected")
}