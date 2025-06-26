package main

import (
	


	_ "example.com/rest-api/docs"                 
	swaggerFiles "github.com/swaggo/files"        
	ginSwagger "github.com/swaggo/gin-swagger"    
	
	"example.com/rest-api/db"
	"example.com/rest-api/middlewares"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

// @title           Event REST API
// @version         1.0
// @description     Simple Gin demo with JWT auth and SQLite storage
// @contact.name    Kaushik
// @contact.url     https://github.com/KaushikML/RestAPI
// @license.name    MIT
// @host            localhost:8080
// @BasePath        /
func main() {
	db.InitDB()
	server := gin.Default()
	server.Use(middlewares.CORSMiddleware())

	routes.RegisterRoutes(server)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8080") // localhost:8080
}