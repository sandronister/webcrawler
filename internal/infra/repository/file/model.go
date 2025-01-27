package file

import (
	"github.com/sandronister/webcrawler/internal/infra/system"
	"github.com/sandronister/webcrawler/internal/ports/iparser"
)

type Model struct {
	folder string
	parser iparser.Type
	system *system.SystemOS
}

func NewFileRepository(folder string, parser iparser.Type, system *system.SystemOS) *Model {
	return &Model{
		folder: folder,
		parser: parser,
		system: system,
	}
}
