package v1

import (
	"eskept/internal/app/context"

	"github.com/gin-gonic/gin"
)

func SetupV1Routes(group *gin.RouterGroup, ctx *context.AppContext) {
	setupAuthGroup(group, ctx)
	setupUserGroup(group, ctx)
	setupProfileGroup(group, ctx)
	setupBusinessGroup(group, ctx)
	setupProviderGroup(group, ctx)
}
