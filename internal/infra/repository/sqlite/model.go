package sqlite

import (
	"database/sql"

	"github.com/sandronister/webcrawler/pkg/logger/types"
)

type Model struct {
	db     *sql.DB
	logger types.ILogger
}

func NewSQLiteRepository(db *sql.DB, logger types.ILogger) (*Model, error) {
	model := &Model{db: db, logger: logger}
	err := model.Create()
	if err != nil {
		return nil, err
	}

	return model, nil
}
