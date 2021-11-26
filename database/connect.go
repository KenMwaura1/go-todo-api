package database

import (
	"fmt"
	"github.com/KenMwaura1/go-todo-api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"strconv"
)

func connectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err = strconv.ParseUint(p, 10, 32)

	configData := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"), port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"))

	DB, err = gorm.Open("postgres", configData)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	fmt.Println("Connected to database")

}
