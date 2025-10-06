package services

import (
	"errors"
	"os"
	"time"

	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/cmd/api/requests"
	"github.com/bermanbenjamin/futStats/internal/logger"
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = os.Getenv("SECRET_KEY")

type Claims struct {
	Username string `json:"username"`
	PlayerId string `json:"player_id"`
	jwt.RegisteredClaims
}

type AuthService struct {
	playerService *PlayerService
	logger        *logger.Logger
}

func NewAuthService(playerService *PlayerService) *AuthService {
	return &AuthService{
		playerService: playerService,
		logger:        logger.GetGlobal().WithComponent("auth_service"),
	}
}

func (s *AuthService) Login(email string, password string) (player *models.Player, token string, err error) {
	s.logger.Info("Login attempt", zap.String("email", email))

	player, err = s.playerService.GetPlayerBy(constants.EMAIL, email)

	if err != nil {
		s.logger.Warn("Login failed - player not found",
			zap.String("email", email),
			zap.Error(err))
		return nil, "", err
	}

	if player == nil {
		s.logger.Warn("Login failed - player not found",
			zap.String("email", email))
		return nil, "", errors.New("player not found")
	}

	if err := checkPassword(password, player.Password); err != nil {
		s.logger.Warn("Login failed - invalid password",
			zap.String("email", email),
			zap.String("player_id", player.ID.String()))
		return nil, "", errors.New("server or password is incorrect")
	}

	token, err = createToken(email, player.ID.String())

	if err != nil {
		s.logger.Error("Login failed - token creation error",
			zap.String("email", email),
			zap.String("player_id", player.ID.String()),
			zap.Error(err))
		return nil, "", err
	}

	s.logger.LogAuthEvent("login", player.ID.String(), email, true)
	return player, token, nil
}

func (s *AuthService) SignUp(request *requests.SignInRequest) (player *models.Player, token string, err error) {
	s.logger.Info("Signup attempt",
		zap.String("email", request.Email),
		zap.String("name", request.Name),
		zap.Int("age", request.Age))

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("Signup failed - password hashing error",
			zap.String("email", request.Email),
			zap.Error(err))
		return nil, "", err
	}

	player = &models.Player{
		Email:    request.Email,
		Password: string(hashedPassword),
		Age:      request.Age,
		Name:     request.Name,
	}

	player, err = s.playerService.CreatePlayer(player)
	if err != nil {
		s.logger.Error("Signup failed - player creation error",
			zap.String("email", request.Email),
			zap.Error(err))
		return nil, "", err
	}

	token, err = createToken(request.Name, player.ID.String())

	if err != nil {
		s.logger.Error("Signup failed - token creation error",
			zap.String("email", request.Email),
			zap.String("player_id", player.ID.String()),
			zap.Error(err))
		return nil, "", err
	}

	s.logger.LogAuthEvent("signup", player.ID.String(), request.Email, true)
	return player, token, nil
}

func (s *AuthService) Logout(tokenString string) (string, error) {
	s.logger.Info("Logout attempt")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		s.logger.Warn("Logout failed - invalid token", zap.Error(err))
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		s.logger.Warn("Logout failed - invalid token claims")
		return "", errors.New("invalid token claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		s.logger.Warn("Logout failed - username not found in claims")
		return "", errors.New("username not found in token claims")
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-5 * time.Minute)),
		},
	})

	tokenString, err = newToken.SignedString([]byte(secretKey))
	if err != nil {
		s.logger.Error("Logout failed - token signing error",
			zap.String("username", username),
			zap.Error(err))
		return "", err
	}

	s.logger.LogAuthEvent("logout", "", username, true)
	return tokenString, nil
}

func checkPassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func createToken(username string, playerId string) (string, error) {
	claims := &Claims{
		Username: username,
		PlayerId: playerId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		// Use global logger since this is a standalone function
		logger.GetGlobal().Error("Token signing failed",
			zap.String("username", username),
			zap.String("player_id", playerId),
			zap.Error(err))
		return "", errors.New("failed to signed token string")
	}

	return tokenString, nil
}
