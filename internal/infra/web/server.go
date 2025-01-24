package web

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sandronister/webcrawler/internal/infra/handler"
)

type Server struct {
	router  *echo.Echo
	webPort string
}

func NewServer(webPort string) *Server {
	return &Server{
		router:  echo.New(),
		webPort: fmt.Sprintf(":%s", webPort),
	}
}

func (s *Server) Start() error {
	return s.router.Start(s.webPort)
}

func (s *Server) AddPageHandler(t *handler.PageHandler) {
	public := s.router.Group("/api/v1")

	public.POST("/start", t.StartCrappint)
}
