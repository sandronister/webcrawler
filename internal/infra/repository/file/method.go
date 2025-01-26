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

	if len(sanitizedUrl) > 100 {
		sanitizedUrl = sanitizedUrl[:100]
	}

	filePath := fmt.Sprintf("%s/%s.txt", r.folder, sanitizedUrl)
	return filePath
}

func (r *Model) Insert(url, content string) error {
	if strings.Contains(url, ".png") || strings.Contains(url, ".css") {
		return nil
	}

	filePath := r.getFilePath(url)

	err := r.system.File(content, filePath)

	if err != nil {
		return fmt.Errorf("error writing file: %s", err.Error())
	}

	return nil
}
