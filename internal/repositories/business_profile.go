package repositories

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/errors"
	"eskept/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusinessProfileRepository struct {
	db *gorm.DB
}

func NewBusinessProfileRepository(ctx *context.AppContext) *BusinessProfileRepository {
	return &BusinessProfileRepository{
		db: ctx.DB,
	}
}

func (bpr *BusinessProfileRepository) Create(profile *models.BusinessProfile) error {
	return bpr.db.Create(profile).Error
}

func (bpr *BusinessProfileRepository) Update(profile *models.BusinessProfile) error {
	return bpr.db.Save(profile).Error
}

func (bpr *BusinessProfileRepository) FindByUserId(userId uuid.UUID) (*models.BusinessProfile, error) {
	var profile models.BusinessProfile
	err := bpr.db.Where("representative_user_id = ?", userId).First(&profile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return &profile, nil
}

func (bpr *BusinessProfileRepository) FindByBusinessProfileId(businessProfileId uuid.UUID) (*models.BusinessProfile, error) {
	var profile models.BusinessProfile
	err := bpr.db.Where("id = ?", businessProfileId).First(&profile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return &profile, nil
}
