package system

import (
	"fmt"
	"os"
)

type SystemOS struct {
}

func NewOS() *SystemOS {
	return &SystemOS{}
}

func (s *SystemOS) Folder(name string) error {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		os.Mkdir(name, 0755)
	}
	return nil
}

func (s *SystemOS) File(content, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("error opening file: %s", err.Error())
	}

	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
		return fmt.Errorf("error writing file: %s", err.Error())
	}

	return nil
}
