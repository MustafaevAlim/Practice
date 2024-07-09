package controllers

import (
	"Practice/internal/model"
	"Practice/internal/repository/user"
	"Practice/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Authorization godoc
// @Summary      Авторизация пользователя
// @Description  Авторизует пользователя и возвращает JWT токен
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "Пользователь"
// @Success 	 200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  string
// @Router       /auth [post]
func (d *Database) Authorization(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	userDB, err := user.GetFromEmail(u.Email, d.DB)
	if err != nil {
		panic(err)
	}

	hashedPassword := userDB.Password
	if !(service.CheckHashPassword(u.Password, hashedPassword)) {
		panic(err)
	}
	token := service.GenerateJWT(u.Email)
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
