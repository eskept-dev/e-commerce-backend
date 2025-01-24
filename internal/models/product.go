package models

import (
	"eskept/internal/utils"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ProductTableName = "products"
)

type Product struct {
	BaseModel

	IsEnabled bool   `gorm:"default:true" json:"isEnabled"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	CodeName  string `gorm:"type:varchar(255)" json:"codeName"`

	Description      string      `gorm:"type:text" json:"description"`
	ShortDescription string      `gorm:"type:text" json:"shortDescription"`
	Details          interface{} `gorm:"type:jsonb" json:"details"`

	ServiceId uuid.UUID `gorm:"type:uuid;not null" json:"serviceId"`
	Service   Service   `gorm:"foreignKey:ServiceId" json:"-"`

	ProviderId uuid.UUID `gorm:"type:uuid;not null" json:"providerId"`
	Provider   Provider  `gorm:"foreignKey:ProviderId" json:"-"`
}

func (p *Product) TableName() string {
	return ProductTableName
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.CodeName = p.GenerateCodeName(p.Name, tx)
	return nil
}

func (p *Product) BeforeUpdate(tx *gorm.DB) error {
	p.CodeName = p.GenerateCodeName(p.Name, tx)
	return nil
}

func (p *Product) GenerateCodeName(name string, tx *gorm.DB) string {
	codeName := utils.ToSnakeCase(name)

	counter := int64(0)
	tx.Where("code_name Like '%"+codeName+"%'", codeName).Count(&counter)

	if counter > 0 {
		codeName = codeName + "_" + strconv.Itoa(int(counter))
	}

	return codeName
}
