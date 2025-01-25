package di

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/internal/infra/handler/page"
	echoserver "github.com/sandronister/webcrawler/internal/infra/web/echo_server"
	"github.com/sandronister/webcrawler/internal/usecase"
)

func NewServer(ports string, broker types.IBroker) *echoserver.Model {
	server := echoserver.NewServer(ports)
	handler := NewPageHandler(broker)
	server.AddPageHandler(handler)
	return server
}

func NewPageHandler(broker types.IBroker) *page.Model {
	return page.NewPageHandler(NewPageUseCase(broker))
}

func NewPageUseCase(broker types.IBroker) *usecase.PageUsecase {
	return usecase.NewPageUsecase(broker)
}
