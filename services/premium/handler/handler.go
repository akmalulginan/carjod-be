package handler

import (
	"net/http"

	"github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/utils"
	"github.com/gin-gonic/gin"
)

type premiumHandler struct {
	premiumUsecase domain.PremiumUsecase
}

func NewPremiumHandler(noAuth, auth *gin.RouterGroup, iuc domain.PremiumUsecase) {
	handler := &premiumHandler{
		premiumUsecase: iuc,
	}
	auth.POST("/premium", handler.Upgrade)
	noAuth.POST("/premium/webhook", handler.Webhook)
}

func (h premiumHandler) Upgrade(ctx *gin.Context) {
	var premium domain.Premium
	err := ctx.Bind(&premium)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to upgrade",
			Error:   err.Error(),
		})
		return
	}

	userId := ctx.GetString(string(utils.KeyUserId))
	if userId == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.Response{
			Message: "failed to upgrade",
			Error:   "invalid user",
		})
		return
	}

	premium.UserId = userId

	err = h.premiumUsecase.Upgrade(ctx, premium)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to upgrade",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, utils.Response{Message: "success"})
}

func (h premiumHandler) Webhook(ctx *gin.Context) {
	var premium domain.Premium
	err := ctx.Bind(&premium)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to upgrade",
			Error:   err.Error(),
		})
		return
	}

	err = h.premiumUsecase.Webhook(ctx, premium)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to upgrade",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, utils.Response{Message: "success"})
}
