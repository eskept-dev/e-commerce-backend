package models

import (
	"eskept/internal/constants/errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return nil
}

func (base *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	base.UpdatedAt = time.Now()
	return nil
}

func (base *BaseModel) Delete(tx *gorm.DB) error {
	if base.DeletedAt != nil || base.IsDeleted {
		return errors.ErrAlreadyDeleted
	}

	now := time.Now()
	base.DeletedAt = &now
	base.IsDeleted = true

	return nil
}
