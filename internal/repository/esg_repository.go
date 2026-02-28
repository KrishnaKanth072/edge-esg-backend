package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/edgeesg/edge-esg-backend/internal/models"
	"gorm.io/gorm"
)

type ESGRepository struct {
	db *gorm.DB
}

func NewESGRepository(db *gorm.DB) *ESGRepository {
	return &ESGRepository{db: db}
}

func (r *ESGRepository) Create(ctx context.Context, score *models.ESGScore) error {
	return r.db.WithContext(ctx).Create(score).Error
}

func (r *ESGRepository) FindByBankID(ctx context.Context, bankID uuid.UUID) ([]models.ESGScore, error) {
	var scores []models.ESGScore
	err := r.db.WithContext(ctx).Where("bank_id = ?", bankID).Find(&scores).Error
	return scores, err
}

func (r *ESGRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.ESGScore, error) {
	var score models.ESGScore
	err := r.db.WithContext(ctx).First(&score, "id = ?", id).Error
	return &score, err
}
