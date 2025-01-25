package page

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sandronister/webcrawler/internal/dto"
)

func (h *Model) StartCrappint(c echo.Context) error {
	pageDTO := new(dto.PageDTO)

	if err := c.Bind(pageDTO); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.GetPage(pageDTO.URL); err != nil {
		return c.JSON(500, map[string]string{"error": "Oops, something went wrong"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Page sent to the broker"})
}
