package file

import (
	"fmt"
	"strings"
)

func (r *Model) getFilePath(url string) string {
	r.system.Folder(r.folder)

	sanitizedUrl := strings.ReplaceAll(url, "://", "_")
	sanitizedUrl = strings.ReplaceAll(sanitizedUrl, "/", "_")
	sanitizedUrl = strings.ReplaceAll(sanitizedUrl, ".", "_")

	filePath := fmt.Sprintf("%s/%s.txt", r.folder, sanitizedUrl)
	return filePath
}

func (r *Model) Insert(url, content string) error {
	if strings.Contains(url, ".png") || strings.Contains(url, ".css") {
		return nil
	}

	textBody := r.parser.GetTagContet(content, "body")

	if len(textBody) == 0 {
		return fmt.Errorf("no body content found")
	}

	contentFile := strings.Join(textBody, "\n")
	filePath := r.getFilePath(url)

	err := r.system.File(contentFile, filePath)

	if err != nil {
		return fmt.Errorf("error writing file: %s", err.Error())
	}

	return nil
}
