package v1

import (
	"Rating-management/internal/models"
	_ "Rating-management/internal/models"
)

type RatingInterface interface {
	FindAll() ([]*models.Rating, error)
	FindById(offerId string, playerId string) (*models.Rating, error)
	Update(rating *models.Rating) error
	Create(rating *models.Rating) (models.Rating, error)
}

type UseCaseInterface interface {
	FindAll() ([]*models.Rating, error)
	FindById(offerId string, playerId string) (*models.Rating, error)
	Update(rating *models.Rating) error
	Create(rating *models.Rating) (models.Rating, error)
}
