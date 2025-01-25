package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"eskept/internal/constants/enums"
	"eskept/internal/utils"
	"strconv"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

const (
	ProductTableName = "products"
)

type AirportTransferDetails struct {
	NumberOfSeats   int `gorm:"type:int" json:"numberOfSeats"`
	NumberOfLuggage int `gorm:"type:int" json:"numberOfLuggage"`
}

func (a *AirportTransferDetails) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, a)
}

func (a AirportTransferDetails) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *AirportTransferDetails) TableName() string {
	return "airport_transfer_details"
}

type FastTrackDetails struct {
	AvailableStartTime string `gorm:"type:varchar(255)" json:"availableStartTime"`
	AvailableEndTime   string `gorm:"type:varchar(255)" json:"availableEndTime"`
}

func (f *FastTrackDetails) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, f)
}

func (f FastTrackDetails) Value() (driver.Value, error) {
	return json.Marshal(f)
}

type EVisaDetails struct {
	ProcessingTime float64 `gorm:"type:float" json:"processingTime"`
}

func (e *EVisaDetails) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, e)
}

func (e EVisaDetails) Value() (driver.Value, error) {
	return json.Marshal(e)
}

type AnalysisMetrics struct {
	TotalBooking int `gorm:"type:int" json:"totalBooking"`
	AvgRating    int `gorm:"type:int" json:"avgRating"`
}

func (a *AnalysisMetrics) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, a)
}

func (a AnalysisMetrics) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type Product struct {
	BaseModel
	IsEnabled bool `gorm:"default:true" json:"isEnabled"`

	Name             string `gorm:"type:varchar(255)" json:"name"`
	CodeName         string `gorm:"type:varchar(255)" json:"codeName"`
	ThumbnailURL     string `gorm:"type:varchar(255)" json:"thumbnailURL"`
	ImageURL         string `gorm:"type:varchar(255)" json:"imageURL"`
	Description      string `gorm:"type:text" json:"description"`
	ShortDescription string `gorm:"type:text" json:"shortDescription"`

	UnitType           enums.ProductUnitType `gorm:"type:varchar(255)" json:"unitType"`
	AnalysisMetrics    AnalysisMetrics       `gorm:"type:jsonb" json:"analysisMetrics"`
	Details            interface{}           `gorm:"type:jsonb" json:"details"`
	Highlights         string                `gorm:"type:text" json:"highlights"`
	CancellationPolicy string                `gorm:"type:text" json:"cancellationPolicy"`

	ServiceID uuid.UUID `gorm:"type:uuid" json:"serviceId"`
	Service   Service   `gorm:"foreignKey:ServiceID" json:"-"`

	ProviderID uuid.UUID `gorm:"type:uuid" json:"providerId"`
	Provider   Provider  `gorm:"foreignKey:ProviderID" json:"-"`
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
