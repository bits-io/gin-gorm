package database

import (
	"gorm.io/gorm"

	"gin_gorm/config"
)

var DB *gorm.DB

func InitDB(cfg *config.AppConfig) {
	DB = config.InitDB(cfg)
}