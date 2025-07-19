package main

import (
	"log"
	"marketplace-api/config"
	"marketplace-api/internal/advert"
	"marketplace-api/internal/auth"
	"marketplace-api/internal/user"
	"marketplace-api/migrations"
	"marketplace-api/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	db := db.NewDb(config)
	migrations.Automigrate(config.Db.Dsn)

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	v1 := router.Group("/api/v1")

	userRepository := user.NewUserRepository(db)
	advertRepository := advert.NewAdvertRepository(db)

	authService := auth.NewAuthService(auth.AuthServiceDeps{
		Config:         config,
		UserRepository: userRepository,
	})

	auth.NewAuthHandler(v1, auth.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})
	advert.NewAdvertHandler(v1, advert.AdvertHandlerDeps{
		Config:           config,
		AdvertRepository: advertRepository,
	})

	if err := router.Run(":" + config.Rest.Port); err != nil {
		log.Fatalf("Failed to start REST API server: %v", err)
	}
}
