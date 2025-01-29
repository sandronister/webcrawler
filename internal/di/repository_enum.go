package di

type RepositoryKind string

const (
	FileRepository   RepositoryKind = "file"
	SQLiteRepository RepositoryKind = "sqlite"
	BadgerRepository RepositoryKind = "badger"
)
