package di

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/internal/infra/handler"
	"github.com/sandronister/webcrawler/internal/infra/web"
	"github.com/sandronister/webcrawler/internal/usecase"
)

func NewServer(ports string, broker types.IBroker) *web.Server {
	server := web.NewServer(ports)
	handler := NewPageHandler(broker)
	server.AddPageHandler(handler)
	return server
}

func NewPageHandler(broker types.IBroker) *handler.PageHandler {
	return handler.NewPageHandler(NewPageUseCase(broker))
}

func NewPageUseCase(broker types.IBroker) *usecase.PageUsecase {
	return usecase.NewPageUsecase(broker)
}
