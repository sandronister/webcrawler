package badger

import (
	"fmt"

	"github.com/dgraph-io/badger/v3"
	"github.com/sandronister/webcrawler/config"
)

type Model struct {
	db *badger.DB
}

func NewBadgerRepository(env *config.Enviroment) (*Model, error) {
	fileBadger := env.RepositoryFile

	if fileBadger == "" {
		return nil, fmt.Errorf("error loading enviroment variables for badger")
	}

	opts := badger.DefaultOptions(fileBadger)
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &Model{db: db}, nil
}
