package main

import (
	"gin-gorm/config"
	"gin-gorm/handler"
	"gin-gorm/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	database := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(database)

	server := gin.Default()

	// middleware CORS
	server.Use(
		middleware.CORSMiddleware(),
	)

	userHandler := handler.UserHandler{DB: database}
	phoneNumberHandler := handler.PhoneNumberHandler{DB: database}

	server.GET("/user", userHandler.HandleGetUser)
	server.POST("/user", userHandler.HandleInsertUser)
	server.GET("/user/:id", userHandler.HandleGetUserByID)

	server.POST("/phone", phoneNumberHandler.HandleInsertPhoneNumber)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
