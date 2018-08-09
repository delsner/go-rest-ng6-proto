package config

import (
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/jinzhu/gorm"
    "fmt"
    "github.com/delsner/go-rest-ng6-proto/backend/models"
    "github.com/delsner/go-rest-ng6-proto/backend/utils"
)

var db *gorm.DB //database

func init() {

    username := utils.GetEnv("db_user", "go_usr")
    password := utils.GetEnv("db_pass", "go_pwd")
    dbName := utils.GetEnv("db_name", "go_db")
    dbHost := utils.GetEnv("db_host", "localhost")

    dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
    fmt.Println(dbUri)

    conn, err := gorm.Open("postgres", dbUri)
    if err != nil {
        fmt.Print(err)
    }

    db = conn
    // Database migration
    db.Debug().AutoMigrate(&models.User{})
    db.Debug().AutoMigrate(&models.Address{})
    db.Debug().AutoMigrate(&models.CreditCard{})
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
    return db
}
