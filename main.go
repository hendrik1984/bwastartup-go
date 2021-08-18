package main

import (
	"bwastartup/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Declare mysql database connection
	dsn := "root:root@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "NameT"
	userInput.Email = "namet@gmail.com"
	userInput.Occupation = "web developer"
	userInput.Password = "namet"

	userService.RegisterUser(userInput)

	// input register dari user
	// handler, mapping input dari user -> struct input
	// service : melakukan mapping dari struct input input ke struct user
	// repository
	// db
}
