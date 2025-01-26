package iparser

type Type interface {
	GetTagContet(input, tag string) []string
	RemoveHTMLTags(input string) string
	ExtractLinks(cContent string, cUrl chan<- string)
}
