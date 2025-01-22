package ports

type IParser interface {
	GetTagContet(input, tag string) []string
	RemoveHTMLTags(input string) string
	ExtractLinks(cContent <-chan string, cUrl chan<- string)
}
