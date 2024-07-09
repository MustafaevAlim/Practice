package controllers

import (
	"Practice/internal/model"
	"Practice/internal/repository/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Registration регистрирует нового пользователя
// @Summary Регистрация нового пользователя
// @Description Регистрирует нового пользователя на основе переданных данных
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "Данные пользователя"
// @Success 200 {object} model.User "Успешная регистрация пользователя"
// @Failure 400 {string} string "Неверный ввод данных"
// @Router /register [post]
func (d *Database) Registration(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	user.InRepo(*u, d.DB)
	return c.JSON(http.StatusOK, u)
}
