package handlers

import (
	"net/http"

	services "github.com/bermanbenjamin/futStats/internal/service"
	"github.com/bermanbenjamin/futStats/internal/transport/http/requests"
	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"username" binding:"required" validate:"required, email"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var credentials Credentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	player, token, err := h.authService.Login(credentials.Username, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"player": player, "token": token})
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("Authorization")

	if tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if len(tokenString) < len("Bearer ") || tokenString[:len("Bearer ")] != "Bearer " {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	token, err := h.authService.Logout(tokenString[len("Bearer "):])
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) SignUp(ctx *gin.Context) {
	var request = requests.SignInRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	player, token, err := h.authService.SignUp(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign in"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "user": player})
}
