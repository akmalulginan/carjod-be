package handler

import (
	"net/http"

	"github.com/akmalulginan/carjod-be/domain"
	"github.com/akmalulginan/carjod-be/utils"
	"github.com/gin-gonic/gin"
)

type matchHandler struct {
	matchUsecase domain.MatchUsecase
}

func NewMatchHandler(r *gin.RouterGroup, iuc domain.MatchUsecase) {
	handler := &matchHandler{
		matchUsecase: iuc,
	}
	r.GET("/matches", handler.GetMatches)
	r.POST("/match", handler.Action)
	r.GET("/match/candidate", handler.GetCandidate)
}

func (h matchHandler) GetMatches(ctx *gin.Context) {

	userId := ctx.GetString(string(utils.KeyUserId))
	if userId == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.Response{
			Message: "failed to get profile",
			Error:   "invalid user",
		})
		return
	}

	users, err := h.matchUsecase.GetMatches(ctx, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to register",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, utils.Response{Message: "success", Data: users})
}

func (h matchHandler) GetCandidate(ctx *gin.Context) {

	userId := ctx.GetString(string(utils.KeyUserId))
	if userId == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.Response{
			Message: "failed to get profile",
			Error:   "invalid user",
		})
		return
	}

	user, err := h.matchUsecase.GetCandidate(ctx, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to register",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, utils.Response{Message: "success", Data: user})
}

func (h matchHandler) Action(ctx *gin.Context) {
	var match domain.Match
	err := ctx.ShouldBindJSON(&match)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to register",
			Error:   err.Error(),
		})
		return
	}
	userId := ctx.GetString(string(utils.KeyUserId))
	if userId == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.Response{
			Message: "failed to get profile",
			Error:   "invalid user",
		})
		return
	}

	match.UserId = userId

	err = h.matchUsecase.Action(ctx, &match)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Message: "failed to register",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, utils.Response{Message: "success", Data: match})
}
