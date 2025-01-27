package handlers

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/enums"
	"eskept/internal/repositories"
	"eskept/internal/services"
	"eskept/internal/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	repo    *repositories.LocationRepository
	service *services.LocationService
	appCtx  *context.AppContext
}

func NewLocationHandler(
	repo *repositories.LocationRepository,
	service *services.LocationService,
	appCtx *context.AppContext,
) *LocationHandler {
	return &LocationHandler{
		repo:    repo,
		service: service,
		appCtx:  appCtx,
	}
}

func (h *LocationHandler) List(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	locationType := enums.LocationType(ctx.Query("locationType"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	pagination := &types.Pagination{
		Page:     page,
		PageSize: pageSize,
	}

	locations, err := h.service.List(keyword, locationType, pagination)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, types.PaginatedResponse{
		Pagination: *pagination,
		Data:       locations,
	})
}
