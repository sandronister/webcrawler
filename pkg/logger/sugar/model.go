package sugar

import "github.com/sandronister/webcrawler/pkg/logger/types"

type Model struct {
	sugar types.ISuggarEngine
}

func NewLogger(sugar types.ISuggarEngine) types.ILogger {
	return &Model{
		sugar: sugar,
	}
}
