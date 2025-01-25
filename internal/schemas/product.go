package schemas

import (
	"eskept/internal/constants/enums"
	"time"

	"github.com/google/uuid"
)

type AirportTransferDetails struct {
	NumberOfSeats   int `json:"numberOfSeats"`
	NumberOfLuggage int `json:"numberOfLuggage"`
}

type FastTrackDetails struct {
	AvailableStartTime string `json:"availableStartTime"`
	AvailableEndTime   string `json:"availableEndTime"`
}

type EVisaDetails struct {
	ProcessingTime float64 `json:"processingTime"`
}

type AnalysisMetrics struct {
	TotalBooking int `json:"totalBooking"`
	AvgRating    int `json:"avgRating"`
}

type ProductCreateRequest struct {
	Name               string                `json:"name" binding:"required"`
	ThumbnailURL       string                `json:"thumbnailURL"`
	ImageURL           string                `json:"imageURL"`
	Description        string                `json:"description"`
	ShortDescription   string                `json:"shortDescription"`
	UnitType           enums.ProductUnitType `json:"unitType" binding:"required"`
	Details            interface{}           `json:"details"`
	Highlights         []string              `json:"highlights"`
	CancellationPolicy string                `json:"cancellationPolicy"`
	ServiceID          uuid.UUID             `json:"serviceId" binding:"required"`
	ProviderID         uuid.UUID             `json:"providerId" binding:"required"`
}

type ProductUpdateRequest struct {
	IsEnabled          bool                  `json:"isEnabled"`
	Name               string                `json:"name" binding:"required"`
	ThumbnailURL       string                `json:"thumbnailURL"`
	ImageURL           string                `json:"imageURL"`
	Description        string                `json:"description"`
	ShortDescription   string                `json:"shortDescription"`
	UnitType           enums.ProductUnitType `json:"unitType" binding:"required"`
	Details            interface{}           `json:"details"`
	Highlights         []string              `json:"highlights"`
	CancellationPolicy string                `json:"cancellationPolicy"`
	ServiceID          uuid.UUID             `json:"serviceId" binding:"required"`
	ProviderID         uuid.UUID             `json:"providerId" binding:"required"`
}

type ProductResponse struct {
	ID                 uuid.UUID             `json:"id"`
	CreatedAt          time.Time             `json:"createdAt"`
	UpdatedAt          time.Time             `json:"updatedAt"`
	IsEnabled          bool                  `json:"isEnabled"`
	Name               string                `json:"name"`
	CodeName           string                `json:"codeName"`
	ThumbnailURL       string                `json:"thumbnailURL"`
	ImageURL           string                `json:"imageURL"`
	Description        string                `json:"description"`
	ShortDescription   string                `json:"shortDescription"`
	UnitType           enums.ProductUnitType `json:"unitType"`
	Details            interface{}           `json:"details"`
	Highlights         []string              `json:"highlights"`
	CancellationPolicy string                `json:"cancellationPolicy"`
	AnalysisMetrics    AnalysisMetrics       `json:"analysisMetrics"`
	Service            ServiceResponse       `json:"service"`
	Provider           ProviderResponse      `json:"provider"`
}
