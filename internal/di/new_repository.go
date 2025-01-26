package di

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/infra/repository/file"
	"github.com/sandronister/webcrawler/internal/infra/repository/sqlite"
	"github.com/sandronister/webcrawler/internal/infra/system"
	"github.com/sandronister/webcrawler/internal/ports/iparser"
	"github.com/sandronister/webcrawler/internal/ports/irepository"
	"github.com/sandronister/webcrawler/pkg/logger/types"
)

func newFileRepository(parser iparser.Type) irepository.Type {
	system := system.NewOS()
	return file.NewFileRepository("output", parser, system)
}

func getConnection(env *config.Enviroment) (*sql.DB, error) {
	fileSQLite := env.SQLiteFile

	if fileSQLite == "" {
		fileSQLite = "crawler.db"
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

	if err != nil {
		return nil, err
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

func newRepository(parser iparser.Type, logger types.ILogger, env *config.Enviroment) (irepository.Type, error) {

	if env.RepositoryKind == "file" || env.RepositoryKind == "" {
		return newFileRepository(parser), nil
	}

	if env.RepositoryKind == "sqlite" {
		return newSQLiteRepository(env, logger)
	}

	return nil, fmt.Errorf("repository type %s not found", env.RepositoryKind)
}
