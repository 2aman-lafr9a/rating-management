package repository

import (
	"Rating-management/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) FindById(offerId string, playerId string) (*models.Rating, error) {
	var rating models.Rating
	err := r.db.Raw("SELECT * FROM ratings WHERE offer_id=? AND player_id=?", offerId, playerId).Scan(&rating).Error
	return &rating, err
}

func (r Repository) FindAll() ([]*models.Rating, error) {
	var ratings []*models.Rating
	err := r.db.Find(&ratings).Error
	return ratings, err
}

func (r Repository) Update(rating *models.Rating) error {
	err := r.db.Save(rating).Error
	return err
}

func (r Repository) Create(rating *models.Rating) (models.Rating, error) {
	err := r.db.Create(rating).Error
	return *rating, err

}
