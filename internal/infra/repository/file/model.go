package file

import (
	"github.com/sandronister/webcrawler/internal/infra/system"
	"github.com/sandronister/webcrawler/internal/ports"
)

type Model struct {
	folder string
	parser ports.IParser
	system *system.SystemOS
}

func NewFileRepository(folder string, parser ports.IParser, system *system.SystemOS) *Model {
	return &Model{
		folder: folder,
		parser: parser,
		system: system,
	}
}
