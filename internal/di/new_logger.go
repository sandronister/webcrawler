package di

import (
	"github.com/sandronister/webcrawler/pkg/logger/factory"
	"github.com/sandronister/webcrawler/pkg/logger/types"
)

func NewLogger(pattern string) (types.ILogger, error) {
	return factory.NewLogger(pattern)
}
