package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


var DB *sql.DB

func InitDB() {
    er := godotenv.Load()
    if er != nil {
        log.Fatal("Error loading .env file")
    }

    // Read env variables
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // Print them out for debugging
    fmt.Println("DB_USER:", dbUser)
    fmt.Println("DB_PASSWORD:", dbPass)
    fmt.Println("DB_HOST:", dbHost)
    fmt.Println("DB_PORT:", dbPort)
    fmt.Println("DB_NAME:", dbName)

    // Build DSN
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
        dbUser, dbPass, dbHost, dbPort, dbName)

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Error:", err)
        panic("Cannot connect to the database")
    }

    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)

    fmt.Println("Successfully connected to the database!!")
}
