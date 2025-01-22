package models

import (
	"github.com/google/uuid"
)

type BusinessProfile struct {
	BaseModel

	BusinessName        string `gorm:"type:varchar(255)" json:"businessName"`
	BusinessTaxId       string `gorm:"type:varchar(255)" json:"businessTaxId"`
	BusinessAddress     string `gorm:"type:varchar(255)" json:"businessAddress"`
	BusinessDialCode    string `gorm:"type:varchar(255)" json:"businessDialCode"`
	BusinessPhoneNumber string `gorm:"type:varchar(255)" json:"businessPhoneNumber"`
	BusinessEmail       string `gorm:"type:varchar(255)" json:"businessEmail"`
	BusinessWebsite     string `gorm:"type:varchar(255)" json:"businessWebsite"`
	BusinessNationality string `gorm:"type:varchar(255)" json:"businessNationality"`

	RepresentativeUserId    uuid.UUID   `gorm:"type:uuid;not null" json:"representativeUserId"`
	RepresentativeProfileId uuid.UUID   `gorm:"type:uuid;not null" json:"representativeProfileId"`
	RepresentativeUser      User        `gorm:"foreignKey:RepresentativeUserId" json:"-"`
	RepresentativeProfile   UserProfile `gorm:"foreignKey:RepresentativeProfileId" json:"-"`
}

func (b *BusinessProfile) TableName() string {
	return "business_profiles"
}
