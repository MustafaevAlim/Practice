package controllers

import (
	"Practice/internal/model"
	"Practice/internal/repository/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

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
func (d *Database) Registration(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	user.InRepo(*u, d.DB)
	return c.JSON(http.StatusOK, u)
}
