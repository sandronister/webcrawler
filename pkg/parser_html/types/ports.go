package types

type IPort interface {
	GetTagContet(input, tag string) []string
	RemoveHTMLTags(input string) string
	ExtractLinks(content, urlFilter string, cUrl chan<- PageDTO)
}
