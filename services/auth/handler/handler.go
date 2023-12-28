package handler

import (
	"net/http"

	"github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/utils"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authUsecase domain.AuthUsecase
}

func NewAuthHandler(r *gin.RouterGroup, iuc domain.AuthUsecase) {
	handler := &authHandler{
		authUsecase: iuc,
	}
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
}

func (h authHandler) Register(ctx *gin.Context) {
	var auth domain.Auth
	err := ctx.Bind(&auth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to register",
			Error:   err.Error(),
		})
		return
	}

	err = h.authUsecase.Register(ctx, auth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to register",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, utils.Response{Message: "success"})
}

func (h authHandler) Login(ctx *gin.Context) {
	var auth domain.Auth
	err := ctx.Bind(&auth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to register",
			Error:   err.Error(),
		})
		return
	}

	token, err := h.authUsecase.Login(ctx, auth)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to register",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, utils.Response{Message: "success", Token: token})
}
