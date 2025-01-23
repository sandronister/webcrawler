package ports

type IRepository interface {
	Insert(url, content string) error
}
