package main

import (
	"TemplateService/database"
	"TemplateService/middlewares"
	"TemplateService/models"
	"TemplateService/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	if err := database.Connect(); err != nil {
		panic("Error connecting to database")
	}

	defer database.Close()

	if err := models.Migrate(); err != nil {
		panic("Error migrating models")
	}

	if os.Getenv("ENV") == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		fmt.Printf("Server is running on port %s\n", os.Getenv("PORT"))
	}

	router := gin.New()

	router.Use(middlewares.Cors())

	router.Use(gin.Recovery())
	if os.Getenv("ENV") == "development" {
		router.Use(gin.Logger())
	}

	router.GET("/health", utils.HealtyResponse)

	router.Run()
}
