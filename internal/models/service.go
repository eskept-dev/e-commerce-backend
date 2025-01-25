package models

import (
	"eskept/internal/constants/enums"
	"eskept/internal/utils"

	"gorm.io/gorm"
)

type Service struct {
	BaseModel
	IsEnabled bool `gorm:"default:true" json:"is_enabled"`

	Name string            `gorm:"type:varchar(255)" json:"name"`
	Code string            `gorm:"type:varchar(255);unique" json:"code"`
	Type enums.ServiceType `gorm:"type:varchar(255)" json:"type"`
}

func (Service) TableName() string {
	return "services"
}

func (s *Service) BeforeCreate(tx *gorm.DB) error {
	s.Code = utils.ToSnakeCase(s.Name)
	return nil
}

func (s *Service) BeforeUpdate(tx *gorm.DB) error {
	s.Code = utils.ToSnakeCase(s.Name)
	return nil
}
