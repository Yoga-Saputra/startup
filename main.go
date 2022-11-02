package main

import (
	"log"
	"startup/handler"
	"startup/users"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:P@ssW0rd@tcp(localhost:3306)/golang_startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("register", userHandler.RegisterUser)
	api.POST("login", userHandler.Login)

	router.Run()

}
