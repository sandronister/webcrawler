package ports

import "github.com/sandronister/webcrawler/internal/infra/handler/page"

type Iserver interface {
	Start() error
	AddPageHandler(t *page.Model)
}
