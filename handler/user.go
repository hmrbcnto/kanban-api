package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hmrbcnto.com/gin-api/entities"
	"hmrbcnto.com/gin-api/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(r *gin.Engine, userUsecase usecase.UserUsecase) {
	handler := &UserHandler{
		userUsecase: userUsecase,
	}
	r.POST("/users", handler.CreateUser)
	r.GET("/users", handler.GetAllUsers)
}

func (h UserHandler) CreateUser(ctx *gin.Context) {
	createUserData := entities.CreateUserRequest{}
	// Serialize body into corresponding request schema
	err := ctx.BindJSON(&createUserData)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := h.userUsecase.CreateUser(&createUserData)
	ctx.JSON(http.StatusAccepted, user)
}

func (h UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.userUsecase.GetAllUsers()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusAccepted, users)
}
