package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"eskept/internal/utils"
	"strconv"

	"gorm.io/gorm"
)

const (
	ProviderTableName = "providers"
)

type BusinessInformation struct {
	PhoneNumber string `gorm:"type:varchar(255)" json:"phoneNumber"`
	Email       string `gorm:"type:varchar(255)" json:"email"`
	Address     string `gorm:"type:varchar(255)" json:"address"`
	Website     string `gorm:"type:varchar(255)" json:"website"`
}

// Scan implements the sql.Scanner interface for BusinessInformation
func (b *BusinessInformation) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &b)
}

// Value implements the driver.Valuer interface for BusinessInformation
func (b BusinessInformation) Value() (driver.Value, error) {
	return json.Marshal(b)
}

type ContactInformation struct {
	FirstName   string `gorm:"type:varchar(255)" json:"firstName"`
	LastName    string `gorm:"type:varchar(255)" json:"lastName"`
	Role        string `gorm:"type:varchar(255)" json:"role"`
	Gender      string `gorm:"type:varchar(255)" json:"gender"`
	PhoneNumber string `gorm:"type:varchar(255)" json:"phoneNumber"`
	Email       string `gorm:"type:varchar(255)" json:"email"`
}

// Scan implements the sql.Scanner interface for ContactInformation
func (c *ContactInformation) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &c)
}

// Value implements the driver.Valuer interface for ContactInformation
func (c ContactInformation) Value() (driver.Value, error) {
	return json.Marshal(c)
}

type Provider struct {
	BaseModel
	IsEnabled bool `gorm:"default:true" json:"isEnabled"`

	Name     string `gorm:"type:varchar(255)" json:"name"`
	CodeName string `gorm:"type:varchar(255);unique" json:"codeName"`

	BusinessInformation BusinessInformation `gorm:"type:jsonb" json:"businessInformation"`
	ContactInformation  ContactInformation  `gorm:"type:jsonb" json:"contactInformation"`
}

func (p *Provider) TableName() string {
	return ProviderTableName
}

func (p *Provider) BeforeCreate(tx *gorm.DB) error {
	p.CodeName = p.GenerateCodeName(p.Name, tx)
	return nil
}

func (p *Provider) BeforeUpdate(tx *gorm.DB) error {
	p.CodeName = p.GenerateCodeName(p.Name, tx)
	return nil
}

func (Provider) GenerateCodeName(name string, tx *gorm.DB) string {
	codeName := utils.ToSnakeCase(name)

	counter := int64(0)
	tx.Where("code_name Like '%"+codeName+"%'", codeName).Count(&counter)

	if counter > 0 {
		codeName = codeName + "_" + strconv.Itoa(int(counter))
	}

	return codeName
}
