package services

import (
	"eskept/internal/app/context"
	"eskept/internal/repositories"
)

type UserService struct {
	repo   *repositories.UserRepository
	appCtx *context.AppContext
}

func NewUserService(
	repo *repositories.UserRepository,
	appCtx *context.AppContext,
) *UserService {
	return &UserService{repo: repo, appCtx: appCtx}
}
