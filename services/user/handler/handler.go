package handler

import (
	"errors"
	"net/http"

	"github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/utils"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(r *gin.RouterGroup, iuc domain.UserUsecase) {
	handler := &userHandler{
		userUsecase: iuc,
	}
	r.GET("/profile", handler.GetProfile)
	r.PUT("/profile", handler.EditProfile)

}

func (h userHandler) GetProfile(ctx *gin.Context) {
	userId := ctx.GetString(string(utils.KeyUserId))
	if userId == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.Response{
			Message: "failed to get profile",
			Error:   "invalid user",
		})
		return
	}

	user, err := h.userUsecase.GetById(ctx, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(400, utils.Response{
			Message: "failed to get profile",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, utils.Response{Message: "success", Data: user})
}

func (h userHandler) EditProfile(ctx *gin.Context) {

	var user domain.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to get profile",
			Error:   err.Error(),
		})
		return
	}

	userId := ctx.GetString(string(utils.KeyUserId))
	if userId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to get profile",
			Error:   errors.New("invalid user id").Error(),
		})
		return
	}

	user.Id = userId
	err = h.userUsecase.Edit(ctx, &user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to get profile",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, utils.Response{Message: "success"})
}
