package main

import (
	"auth-service/config"
	"auth-service/internal/handler"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	db := db.NewDb(config)

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	userRepository := repository.NewUserRepository(db)

	authService := service.NewAuthService(service.AuthServiceDeps{
		Config:         config,
		UserRepository: userRepository,
	})

	handler.NewAuthHandler(router, handler.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})

	if err := router.Run(":" + config.Rest.Port); err != nil {
		log.Fatalf("Failed to start REST API server: %v", err)
	}
}
