package controllers

import (
	"Practice/internal/config"
	"Practice/internal/model"
	"Practice/internal/repository/sales"
	"Practice/internal/repository/user"
	"Practice/internal/service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for a pet store.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide your JWT token as: Bearer {token}
// @security BearerAuth

// Registration godoc
// @Summary      Registers a new user
// @Description  Handles user registration by binding and validating the user input
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "User to register"
// @Success      200   {object}  model.User
// @Failure      400   {string}  string "Bad Request"
// @Router       /register [post]
func Registration(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	user.AddUserInRepo(*u)
	return c.JSON(http.StatusOK, u)
}

// Authorization godoc
// @Summary      Авторизация пользователя
// @Description  Авторизует пользователя и возвращает JWT токен
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "Пользователь"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  string
// @Router       /auth [post]
func Authorization(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	userDB, err := user.GetUserEmail(u.Email)
	if err != nil {
		panic(err)
	}
	hashedPassword := userDB.Password

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password)); err != nil {
		return c.String(http.StatusOK, "Wrong password")
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &user.Claims{
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

// Welcome godoc
// @Summary      Приветствие пользователя
// @Description  Возвращает приветственное сообщение для авторизованного пользователя
// @Tags         user
// @Accept       json
// @Produce      plain
// @Success      200  {string}  string  "Welcome {email}!"
// @Failure      400  {string}  string  "Bad Request"
// @Failure      401  {string}  string  "Unauthorized"
// @Router       /welcome [get]
// @security BearerAuth
func Welcome(c echo.Context) error {
	// Получение email из контекста
	email := c.Get("userEmail").(string)
	return c.String(http.StatusOK, fmt.Sprintf("Welcome %s!", email))
}

// NewSale godoc
// @Summary      Создание новой продажи
// @Description  Создает новую запись о продаже в системе
// @Tags         sales
// @Accept       json
// @Produce      json
// @Param        sale  body    model.Sale  true  "Информация о продаже"
// @Success      200   {object}  model.Sale  "Успешное создание продажи"
// @Failure      400   {object}  map[string]string  "Ошибка в запросе"
// @Router       /sales [post]
// @security BearerAuth
func NewSale(c echo.Context) error {
	s := new(model.Sale)
	if err := c.Bind(s); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	sales.AddSalesInRepo(*s)
	return c.JSON(http.StatusOK, s)
}

// GetPdfReport godoc
// @Summary      Получить PDF отчет
// @Description  Генерирует и возвращает PDF отчет
// @Tags         reports
// @Produce      application/pdf
// @Success      200  {file}  file
// @Failure      500  {object}  map[string]string  "Внутренняя ошибка сервера"
// @Router       /pdfReport [get]
func GetPdfReport(c echo.Context) error {
	service.MakePdf()
	filePath := "/home/traktor/GolandProjects/Practice/internal/service/table.pdf"
	return c.Attachment(filePath, "downloaded_file.pdf")
}

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
