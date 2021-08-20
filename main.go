package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
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

	// di ruby on rails service itu model kumpuluan business logic
	userService := user.NewService(userRepository)

	// di ruby on rails handler itu controlernya yang hubungin ke model yang berkaitan
	userHandler := handler.NewUserHandler(userService)

	// define router, kalau di ruby on rails di taruh di file routes.rb
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "NameT"
	// userInput.Email = "namet@gmail.com"
	// userInput.Occupation = "web developer"
	// userInput.Password = "namet"
	// userService.RegisterUser(userInput)

	// input register dari user
	// handler, mapping input dari user -> struct input
	// service : melakukan mapping dari struct input input ke struct user
	// repository
	// db
}
