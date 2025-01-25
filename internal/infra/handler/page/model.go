package page

import "github.com/sandronister/webcrawler/internal/usecase"

type Model struct {
	usecase *usecase.PageUsecase
}

func NewPageHandler(usecase *usecase.PageUsecase) *Model {
	return &Model{usecase: usecase}
}
