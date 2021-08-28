package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/refaldyrk/auth-go/auth"
	"github.com/refaldyrk/auth-go/handler"
	"github.com/refaldyrk/auth-go/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/auth-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api")

	//USER
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)

	router.Run()
}
