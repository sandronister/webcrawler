package parser

import "regexp"

type HtmlParser struct {
}

func NewHtmlParser() *HtmlParser {
	return &HtmlParser{}
}

func (p *HtmlParser) GetTagContet(input, tag string) []string {
	re := regexp.MustCompile(`<` + tag + `[^>]*>(.*?)</` + tag + `>`)
	matches := re.FindAllStringSubmatch(input, -1)
	var contents []string
	for _, match := range matches {
		if len(match) > 1 {
			contents = append(contents, match[1])
		}
	}
	return contents
}

func (p *HtmlParser) RemoveHTMLTags(input string) string {
	re := regexp.MustCompile(`<.*?>`)
	return re.ReplaceAllString(input, "")
}

func (p *HtmlParser) ExtractLinks(content string, cUrl chan<- string) {
	re := regexp.MustCompile(`href="(http[s]?://[^"]+)"`)
	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) > 1 {
			cUrl <- match[1]
		}
	}
}
