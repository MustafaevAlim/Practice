package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

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
func (d *Database) Welcome(c echo.Context) error {
	email := c.Get("userEmail").(string)
	return c.String(http.StatusOK, fmt.Sprintf("Welcome %s!", email))
}
