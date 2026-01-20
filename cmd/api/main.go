package main

import (
	"log"

	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
	"github.com/GuilhermePT1/api-social-meli/internal/infra/database"
	"github.com/GuilhermePT1/api-social-meli/internal/infra/http/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/GuilhermePT1/api-social-meli/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title SocialMeli API
// @version 1.0
// @description API do projeto SocialMeli
// @host localhost:8080
// @BasePath /
func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error())
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Post{},
		&models.Follow{})

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(router, db)

	router.Run(":8080")
}
