package handlers

import (
	"net/http"

	"github.com/bermanbenjamin/futStats/cmd/api/requests"
	"github.com/bermanbenjamin/futStats/internal/logger"
	"github.com/bermanbenjamin/futStats/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthHandler struct {
	authService *services.AuthService
	logger      *logger.Logger
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger:      logger.GetGlobal().WithComponent("auth_handler"),
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var credentials Credentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		h.logger.Warn("Invalid login credentials format",
			zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	h.logger.Info("Login attempt initiated",
		zap.String("username", credentials.Username))

	player, token, err := h.authService.Login(credentials.Username, credentials.Password)
	if err != nil {
		h.logger.Warn("Login failed",
			zap.String("username", credentials.Username),
			zap.Error(err))
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	h.logger.Info("Login successful",
		zap.String("username", credentials.Username),
		zap.String("player_id", player.ID.String()))

	ctx.JSON(http.StatusOK, gin.H{"player": player, "token": token})
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("Authorization")

	if tokenString == "" {
		h.logger.Warn("Logout attempted without Authorization header")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if len(tokenString) < len("Bearer ") || tokenString[:len("Bearer ")] != "Bearer " {
		tokenPrefix := tokenString
		if len(tokenString) > 20 {
			tokenPrefix = tokenString[:20]
		}
		h.logger.Warn("Logout attempted with invalid token format",
			zap.String("token_prefix", tokenPrefix))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	h.logger.Info("Logout attempt initiated")

	token, err := h.authService.Logout(tokenString[len("Bearer "):])
	if err != nil {
		h.logger.Error("Logout failed",
			zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	h.logger.Info("Logout successful")
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) SignUp(ctx *gin.Context) {
	var request = requests.SignInRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.logger.Warn("Invalid signup request format",
			zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	h.logger.Info("Signup attempt initiated",
		zap.String("email", request.Email),
		zap.String("name", request.Name))

	player, token, err := h.authService.SignUp(&request)
	if err != nil {
		h.logger.Error("Signup failed",
			zap.String("email", request.Email),
			zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign up"})
		return
	}

	h.logger.Info("Signup successful",
		zap.String("email", request.Email),
		zap.String("player_id", player.ID.String()))

	ctx.JSON(http.StatusOK, gin.H{"token": token, "user": player})
}
