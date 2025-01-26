package filter

import (
	"strings"
)

func (f *URLFilter) Filter(match []string, filter string) bool {
	if len(match) < 1 {
		return false
	}

	if filter == "" {
		return true
	}

	return contains(match[1], filter)
}

func contains(url, filter string) bool {
	return strings.Contains(url, filter)
}
