package page

import "github.com/sandronister/webcrawler/internal/usecase/page"

type Model struct {
	usecase *page.Model
}

func NewPageHandler(usecase *page.Model) *Model {
	return &Model{usecase: usecase}
}
