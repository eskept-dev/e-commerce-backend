package models

import (
	"eskept/internal/constants/enums"
	"eskept/internal/utils"

	"gorm.io/gorm"
)

type Location struct {
	BaseModel
	Name           string             `gorm:"type:varchar(255)" json:"name"`
	CodeName       string             `gorm:"type:varchar(255)" json:"codeName"`
	Type           enums.LocationType `gorm:"type:varchar(255)" json:"type"`
	ParentCodeName string             `gorm:"type:varchar(255)" json:"parentCodeName"`
	Priority       int                `gorm:"type:int" json:"priority"`
}

func (l *Location) TableName() string {
	return "locations"
}

func (l *Location) BeforeCreate(tx *gorm.DB) error {
	l.CodeName = l.GenerateCodeName(l.Name, tx)
	return nil
}

func (l *Location) BeforeUpdate(tx *gorm.DB) error {
	l.CodeName = l.GenerateCodeName(l.Name, tx)
	return nil
}

func (l *Location) GenerateCodeName(name string, tx *gorm.DB) string {
	return utils.ToSnakeCase(name)
}
