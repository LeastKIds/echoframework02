package handlers

import (
	"app/cmd/api/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func PostIndexHandler(c echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process data")
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	// Atoi = 문자열을 정수로 바꾸는 것
	idx, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process data")
	}

	data, err := service.GetById(idx)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process data")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}
