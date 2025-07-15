package handler

import (
	"auth-service/config"
	"auth-service/internal/service"
	"auth-service/pkg/jwt"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandlerDeps struct {
	*config.Config
	*service.AuthService
}

type AuthHandler struct {
	*config.Config
	*service.AuthService
}

func NewAuthHandler(router *gin.Engine, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.POST("/auth/register", handler.Register)
	router.POST("/auth/login", handler.Login)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	username, err := h.AuthService.Register(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	token, err := jwt.NewJWT(h.Config.Auth.Secret).Create(jwt.JWTData{
		Username: username,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	username, err := h.AuthService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	token, err := jwt.NewJWT(h.Config.Auth.Secret).Create(jwt.JWTData{
		Username: username,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
