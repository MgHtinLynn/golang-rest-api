package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/spf13/viper"
)

func SetupModels() *gorm.DB {
	viper.AutomaticEnv()

	// To get the value from the config file using key
	viper_user := "postgres"
	viper_password := "admin1234"
	viper_db := "muledb"
	viper_host := "127.0.0.1"
	viper_port := "5432"

	//https://gobyexample.com/string-formatting
	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)
	fmt.Println("conname is\t\t", prosgret_conname)

	db, err := gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&User{})

	// Initialise value
	m := User{Name: "Htin Lynn",Email: "htinlin01@gmail.com",Phone: "09785360975", Address: "Home"}
	db.Create(&m)
	return db
}
