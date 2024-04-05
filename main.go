package main

import (
	"TemplateService/database"
	"TemplateService/middlewares"
	"TemplateService/utils"
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

	utils.InitTranslate()

	if os.Getenv("ENV") == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
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
