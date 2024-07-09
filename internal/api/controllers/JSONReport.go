package controllers

import (
	"Practice/internal/repository/sales"
	"github.com/labstack/echo/v4"
	"net/http"
)

// JSONReport выводит все продажи в формате JSON
// @Summary Получить все продажи
// @Description Получает список всех продаж из базы данных и возвращает его в формате JSON
// @Tags reports
// @Accept  json
// @Produce  json
// @Success 200 {array} model.InfoSales "Список всех продаж"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /JSONReport [get]
func (d *Database) JSONReport(c echo.Context) error {
	sale := sales.GetAllSales(d.DB)
	return c.JSON(http.StatusOK, sale)
}
