package ports

type IRepository interface {
	Insert(content <-chan string) error
}
