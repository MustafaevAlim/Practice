package controllers

import (
	"Practice/internal/repository/sales"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// DelSale удаляет запись о продаже по ID
// @Summary Удаление записи о продаже
// @Description Удаляет запись о продаже по ID
// @Tags sales
// @Accept json
// @Produce json
// @Param id path int true "ID записи о продаже"
// @Success 200 {string} string "Succesfully deleted"
// @Failure 400 {string} string "Неверный ID"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /sales/{id} [delete]
// @security BearerAuth
func (d *Database) DelSale(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	sales.Del(id, d.DB)
	return c.String(http.StatusOK, "Succesfully deleted")
}
