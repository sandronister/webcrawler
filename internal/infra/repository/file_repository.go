package repository

import (
	"fmt"

	"github.com/sandronister/webcrawler/internal/ports"
)

type FileRepository struct {
	folder string
	parser ports.IParser
}

func NewFileRepository(folder string, parser ports.IParser) *FileRepository {
	return &FileRepository{
		folder: folder,
		parser: parser,
	}
}

func (r *FileRepository) Insert(content <-chan string) error {
	for c := range content {
		textBody := r.parser.GetTagContet(c, "body")

		fmt.Print(len(textBody))

	}
	return nil
}
