package services

import (
	"gorm.io/gorm"
	"rekaGo/models"
)

type ImageService interface {
	GetImage(id int64) (models.Image, error)
}

type imageServiceImpl struct {
	db *gorm.DB
}

func NewImageService(db *gorm.DB) ImageService {
	return &imageServiceImpl{
		db: db,
	}
}

func (s *imageServiceImpl) GetImage(id int64) (models.Image, error) {
	var image models.Image

	result := s.db.
		Where("id = ?", id).
		Find(&image)
	return image, result.Error
}
