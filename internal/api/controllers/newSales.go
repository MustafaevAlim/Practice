package controllers

import (
	"Practice/internal/model"
	"Practice/internal/repository/sales"
	"github.com/labstack/echo/v4"
	"net/http"
)

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
	sales.InRepo(*s)
	return c.JSON(http.StatusOK, s)
}
