package repositories

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/errors"
	"eskept/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfileRepository struct {
	db *gorm.DB
}

func NewUserProfileRepository(ctx *context.AppContext) *UserProfileRepository {
	return &UserProfileRepository{db: ctx.DB}
}

func (r *UserProfileRepository) Create(profile *models.UserProfile) error {
	return r.db.Create(profile).Error
}

func (r *UserProfileRepository) FindByUserId(userId uuid.UUID) (*models.UserProfile, error) {
	var profile models.UserProfile
	err := r.db.Where("user_id = ?", userId).First(&profile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserProfileNotFound
		}
		return nil, err
	}
	return &profile, nil
}

func (r *UserProfileRepository) Update(profile *models.UserProfile) error {
	return r.db.Save(profile).Error
}
