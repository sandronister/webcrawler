package html

import "regexp"

func (p *Model) GetTagContet(input, tag string) []string {
	re := regexp.MustCompile(`(?s)<` + tag + `[^>]*>(.*?)</` + tag + `>`)
	matches := re.FindAllStringSubmatch(input, -1)
	var contents []string
	for _, match := range matches {
		if len(match) > 1 {
			contents = append(contents, match[1])
		}
	}
	return contents
}

func (p *Model) RemoveHTMLTags(input string) string {
	re := regexp.MustCompile(`<.*?>`)
	return re.ReplaceAllString(input, "")
}

func (p *Model) ExtractLinks(content string, cUrl chan<- string) {
	re := regexp.MustCompile(`href="(http[s]?://[^"]+)"`)
	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) > 1 {
			cUrl <- match[1]
		}
	}
}
