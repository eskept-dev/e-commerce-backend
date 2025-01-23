package schemas

import (
	"time"

	"github.com/google/uuid"
)

type BusinessProfileCreateRequest struct {
	BusinessName        string `json:"businessName" binding:"required"`
	BusinessTaxId       string `json:"businessTaxId" binding:"required"`
	BusinessAddress     string `json:"businessAddress" binding:"required"`
	BusinessDialCode    string `json:"businessDialCode" binding:"required"`
	BusinessPhoneNumber string `json:"businessPhoneNumber" binding:"required"`
	BusinessEmail       string `json:"businessEmail" binding:"required"`
	BusinessWebsite     string `json:"businessWebsite" binding:"required"`
	BusinessNationality string `json:"businessNationality" binding:"required"`
}

type BusinessProfileCreateResponse struct {
	ID                      uuid.UUID `json:"id"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
	BusinessName            string    `json:"businessName"`
	BusinessTaxId           string    `json:"businessTaxId"`
	BusinessAddress         string    `json:"businessAddress"`
	BusinessDialCode        string    `json:"businessDialCode"`
	BusinessPhoneNumber     string    `json:"businessPhoneNumber"`
	BusinessEmail           string    `json:"businessEmail"`
	BusinessWebsite         string    `json:"businessWebsite"`
	BusinessNationality     string    `json:"businessNationality"`
	RepresentativeUserId    uuid.UUID `json:"representativeUserId"`
	RepresentativeProfileId uuid.UUID `json:"representativeProfileId"`
}

type BusinessProfileGetResponse struct {
	ID                      uuid.UUID `json:"id"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
	BusinessName            string    `json:"businessName"`
	BusinessTaxId           string    `json:"businessTaxId"`
	BusinessAddress         string    `json:"businessAddress"`
	BusinessDialCode        string    `json:"businessDialCode"`
	BusinessPhoneNumber     string    `json:"businessPhoneNumber"`
	BusinessEmail           string    `json:"businessEmail"`
	BusinessWebsite         string    `json:"businessWebsite"`
	BusinessNationality     string    `json:"businessNationality"`
	RepresentativeUserId    uuid.UUID `json:"representativeUserId"`
	RepresentativeProfileId uuid.UUID `json:"representativeProfileId"`
}
