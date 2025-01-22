package repository

import "fmt"

type FileRepository struct {
	folder string
}

func NewFileRepository(folder string) *FileRepository {
	return &FileRepository{
		folder: folder,
	}
}

func (r *FileRepository) Insert(content <-chan string) error {
	for c := range content {
		fmt.Println(c)
	}
	return nil
}
