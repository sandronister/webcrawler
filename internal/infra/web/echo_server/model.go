package echoserver

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Model struct {
	router  *echo.Echo
	webPort string
}

func NewServer(webPort string) *Model {
	return &Model{
		router:  echo.New(),
		webPort: fmt.Sprintf(":%s", webPort),
	}
}
