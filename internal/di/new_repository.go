package di

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/infra/repository/badger"
	"github.com/sandronister/webcrawler/internal/infra/repository/file"
	"github.com/sandronister/webcrawler/internal/infra/repository/sqlite"
	"github.com/sandronister/webcrawler/internal/infra/system"
	"github.com/sandronister/webcrawler/internal/ports/irepository"
	"github.com/sandronister/webcrawler/pkg/logger/types"
	parserType "github.com/sandronister/webcrawler/pkg/parser_html/types"
)

func newFileRepository(parser parserType.IPort) irepository.Type {
	system := system.NewOS()
	return file.NewFileRepository("output", parser, system)
}

func getConnection(env *config.Enviroment) (*sql.DB, error) {
	fileSQLite := env.RepositoryFile

	if fileSQLite == "" {
		return nil, fmt.Errorf("repository file not found")
	}

	conn, err := sql.Open("sqlite3", fileSQLite)

	if err != nil {
		return nil, err
	}

	var result string
	err = conn.QueryRow("PRAGMA integrity_check").Scan(&result)

	if err != nil {
		return nil, err
	}

	if result != "ok" {
		return nil, fmt.Errorf("integrity check failed: %s", result)
	}

	return conn, nil
}

func newSQLiteRepository(env *config.Enviroment, logger types.ILogger) (irepository.Type, error) {
	conn, err := getConnection(env)

	if err != nil {
		return nil, err
	}

	sqlite, err := sqlite.NewSQLiteRepository(conn, logger)
	if err != nil {
		return nil, err
	}

	return sqlite, nil
}

func NewBadgerRepository(env *config.Enviroment) (irepository.Type, error) {
	return badger.NewBadgerRepository(env)
}

func newRepository(parser parserType.IPort, logger types.ILogger, env *config.Enviroment) (irepository.Type, error) {
	switch RepositoryKind(env.RepositoryKind) {
	case FileRepository, "":
		return newFileRepository(parser), nil
	case SQLiteRepository:
		return newSQLiteRepository(env, logger)
	case BadgerRepository:
		return NewBadgerRepository(env)
	default:
		return nil, fmt.Errorf("repository type %s not found", env.RepositoryKind)
	}
}
