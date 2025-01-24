package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sandronister/webcrawler/internal/dto"
	"github.com/sandronister/webcrawler/internal/usecase"
)

type PageHandler struct {
	usecase *usecase.PageUsecase
}

func NewPageHandler(usecase *usecase.PageUsecase) *PageHandler {
	return &PageHandler{usecase: usecase}
}

func (h *PageHandler) StartCrappint(c echo.Context) error {
	pageDTO := new(dto.PageDTO)

	if err := c.Bind(pageDTO); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.GetPage(pageDTO.URL); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Page sent to the broker"})
}
