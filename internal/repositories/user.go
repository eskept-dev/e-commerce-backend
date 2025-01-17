package repositories

import (
	"eskept/internal/models"
	"eskept/pkg/context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewUserRepository(ctx *context.AppContext) *UserRepository {
	return &UserRepository{
		DB: ctx.DB,
	}
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) IsEmailExists(email string) bool {
	var user models.User
	r.DB.First(&user, "email = ?", email)
	return user != (models.User{})
}

func (r *UserRepository) FindByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	return r.DB.Save(user).Error
}
