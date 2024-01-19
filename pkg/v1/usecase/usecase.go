package usecase

import (
	"Rating-management/internal/models"
	interfaces "Rating-management/pkg/v1"
)

type UseCase struct {
	repo interfaces.RatingInterface
}

func (u UseCase) FindAll() ([]*models.Rating, error) {
	ratings, err := u.repo.FindAll()
	return ratings, err
}

func (u UseCase) FindById(offerId string, playerId string) (*models.Rating, error) {
	rating, err := u.repo.FindById(offerId, playerId)
	return rating, err
}

func (u UseCase) Update(rating *models.Rating) error {
	err := u.repo.Update(rating)
	return err
}

func (u UseCase) Create(rating *models.Rating) (models.Rating, error) {
	rat, err := u.repo.Create(rating)
	return rat, err
}

func New(repo interfaces.RatingInterface) interfaces.UseCaseInterface {
	return &UseCase{repo: repo}
}
