
package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database")
    }
}
