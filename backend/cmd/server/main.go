package main

import (
	"backend/internals/users"
	"backend/pkg/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := gin.Default()

	db := database.Connect(os.Getenv("dbURL"))

	userRepo := users.NewRepository(db)
	userService := users.NewService(userRepo)
	userHandler := users.NewHandler(userService)

	users.RegisterRoutes(server, userHandler)

	server.Run(":8080")
}
