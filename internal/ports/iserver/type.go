package iserver

import "github.com/sandronister/webcrawler/internal/infra/handler/page"

type Type interface {
	Start() error
	AddPageHandler(t *page.Model)
}
