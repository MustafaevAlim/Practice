package service

import (
	"Practice/internal/config"
	"Practice/internal/repository/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

// Проверка авторизации
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "missing or malformed jwt"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "missing or malformed jwt"})
		}

		token, err := jwt.ParseWithClaims(tokenStr, &user.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return config.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid jwt signature"})
			}
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
		}

		if claims, ok := token.Claims.(*user.Claims); ok && token.Valid {
			c.Set("userEmail", claims.Email)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized"})
	}
}

func CheckHashPassword(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}

func GenerateJWT(email string) string {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &user.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		return err.Error()
	}
	return tokenString
}
