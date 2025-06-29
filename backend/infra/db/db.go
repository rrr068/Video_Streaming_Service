package db

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
)

var DB *gorm.DB

func InitDB() {
    dsn := os.Getenv("DB_DSN")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(fmt.Sprintf("failed to connect to db: %v", err))
    }
    DB = db
}
