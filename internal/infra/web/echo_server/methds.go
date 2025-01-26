package echoserver

import "github.com/sandronister/webcrawler/internal/infra/handler/page"

func (s *Model) Start() error {
	return s.router.Start(s.webPort)
}

func (s *Model) AddPageHandler(t *page.Model) {
	public := s.router.Group("/api/v1")

	public.POST("/start", t.StartCrappint)
}
