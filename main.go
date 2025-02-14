package main

import (
	"gin_gorm/config"
	"gin_gorm/infra/database"
	"gin_gorm/infra/logger"
	"gin_gorm/routers"
)

func main() {
	// Load configuration
	config.LoadEnv()
	cfg := config.LoadConfig()

	// Initialize logger
	logger.InitLogger()

	// Initialize database
	database.InitDB(cfg)

	// Initialize router
	r := routers.InitRouter()

	// Start server
	r.Run(":8080")
}