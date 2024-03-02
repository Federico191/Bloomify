package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectIntern/internal/model"
	"projectIntern/internal/usecase"
	"projectIntern/pkg/response"
)

type AuthHandler struct {
	userUC usecase.AuthUseCaseItf
}

func NewAuthHandler(userUC usecase.AuthUseCaseItf) *AuthHandler {
	return &AuthHandler{userUC: userUC}
}

func (a AuthHandler) Register(c *gin.Context) {
	var req model.UserRegister

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind request", err)
		return
	}

	user, err := a.userUC.Register(c, &req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	response.Success(c, http.StatusOK, "Success Register user", user)
}

func (a AuthHandler) Login(c *gin.Context) {
	var req model.UserLogin

	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to bind request", err)
	}

	token, err := a.userUC.Login(c, &req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to log in", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": token, "Message": "Success Login"})
}
