package repository

import (
	"gin_gorm/infra/database"
	"gin_gorm/models"
)

func GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	if err := database.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}