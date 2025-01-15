package services

import (
	"errors"
	"log"
	"os"
	"time"

	model "github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/transport/http/constants"
	"github.com/bermanbenjamin/futStats/internal/transport/http/requests"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = os.Getenv("SECRET_KEY")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AuthService struct {
	playerService *PlayerService
}

func NewAuthService(playerService *PlayerService) *AuthService {
	return &AuthService{playerService: playerService}
}

func (s *AuthService) Login(email string, password string) (player *model.Player, token string, err error) {
	player, err = s.playerService.GetPlayerBy(constants.EMAIL, email)

	if err != nil {
		log.Printf("error getting player from server")
		return nil, "", errors.New("error getting player from server with this email: " + email)
	}

	if err := checkPassword(password, player.Password); err != nil {
		log.Printf("error checking password")
		return nil, "", errors.New("invalid password")
	}

	token, err = createToken(email)

	if err != nil {
		return nil, "", err
	}

	return player, token, nil
}

func (s *AuthService) SignUp(request *requests.SignInRequest) (player *model.Player, token string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	player = &model.Player{
		Email:    request.Email,
		Password: string(hashedPassword),
		Age:      request.Age,
		Name:     request.Name,
	}

	player, err = s.playerService.CreatePlayer(player)
	if err != nil {
		return nil, "", err
	}

	token, err = createToken(request.Name)

	if err != nil {
		return nil, "", err
	}

	return player, token, nil
}

func (s *AuthService) Logout(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: claims["username"].(string),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-5 * time.Minute)),
		},
	})

	tokenString, err = newToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func checkPassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func createToken(username string) (string, error) {

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Printf("error signing token: %v", err)
		return "", errors.New("failed to signed token string")
	}

	return tokenString, nil
}
