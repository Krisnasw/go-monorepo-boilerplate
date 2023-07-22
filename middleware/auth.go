package middleware

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"

	"go-monorepo-boilerplate/services/auth-svc/config"
)

var cfg = config.New()
var secretKey = []byte(cfg.JwtSecret)

type Claims struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	jwt.RegisteredClaims
}

// Generate Jwt based on Claims.
func GenerateJwt(id, name, username, email, avatar string) (string, error) {
	expirationTime := time.Now().Add(720 * time.Hour)

	claims := &Claims{
		Id:       id,
		Name:     name,
		Username: username,
		Email:    email,
		Avatar:   avatar,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Verify JWT Token
func VerifyJwt(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		err = errors.New("Token is invalid")
		return
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("Token Expired")
		return
	}

	return
}

// Get Claims Detail
func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, err
	} else {
		logrus.Errorln("Invalid JWT Token")
		return nil, err
	}
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"success": false, "message": "Unauthorized Access"})
			context.Abort()
			return
		}

		err := VerifyJwt(strings.Split(tokenString, "Bearer ")[1])
		if err != nil {
			context.JSON(401, gin.H{"success": false, "message": "Invalid Access Token"})
			context.Abort()
			return
		}
		context.Next()
	}
}
