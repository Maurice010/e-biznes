package database

import (
    "task4/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
    if err != nil {
        panic(err.Error)
    }

    db.AutoMigrate(
        &models.Product{},
        &models.Category{},
        &models.Cart{},
        &models.CartItem{},
		&models.Supplier{},
    )

    DB = db
}