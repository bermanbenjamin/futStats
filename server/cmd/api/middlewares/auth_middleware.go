package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = os.Getenv("SECRET_KEY")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func AuthMiddleware(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("Authorization")

	if tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if len(tokenString) < len("Bearer ") || tokenString[:len("Bearer ")] != "Bearer " {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	tokenString = tokenString[len("Bearer "):]

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(jwtKey), nil
	})

	if err != nil || !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	ctx.Set("username", claims.Username)

	ctx.Next()

}
