package page

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sandronister/webcrawler/pkg/parser_html/types"
)

func (h *Model) StartCrappint(c echo.Context) error {
	pageDTO := new(types.PageDTO)

	if err := c.Bind(pageDTO); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.GetPage(pageDTO); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "URl already started to be crawled"})
}
