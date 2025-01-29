package file

import (
	"github.com/sandronister/webcrawler/internal/infra/system"
	"github.com/sandronister/webcrawler/pkg/parser_html/types"
)

type Model struct {
	folder string
	parser types.IPort
	system *system.SystemOS
}

func NewFileRepository(folder string, parser types.IPort, system *system.SystemOS) *Model {
	return &Model{
		folder: folder,
		parser: parser,
		system: system,
	}
}
