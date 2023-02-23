package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hmrbcnto.com/gin-api/entities"
	"hmrbcnto.com/gin-api/usecase"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(r *gin.Engine, authUsecase usecase.AuthUsecase) {
	handler := &AuthHandler{
		authUsecase: authUsecase,
	}

	r.POST("/login", handler.Login)
}

func (h AuthHandler) Login(ctx *gin.Context) {
	loginBody := entities.Login{}
	// Serialize body into corresponding request schema
	err := ctx.BindJSON(&loginBody)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := h.authUsecase.Login(loginBody.Email, loginBody.Password)
	ctx.JSON(http.StatusAccepted, token)
}
