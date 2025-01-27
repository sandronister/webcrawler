package irepository

type Type interface {
	Insert(url, content string) error
}
