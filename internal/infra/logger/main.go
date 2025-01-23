package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/sandronister/webcrawler/internal/infra/system"
)

type logDTO struct {
	Type      string `json:"type"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	ClassName string `json:"class_name"`
}

type Log struct {
	currentTime time.Time
	system      *system.SystemOS
}

func NewLog(system *system.SystemOS) *Log {
	return &Log{
		system: system,
	}
}

func (l *Log) Info(className, message string) {
	l.currentTime = time.Now()
	logMessage := logDTO{
		Message:   message,
		Timestamp: l.currentTime.Format("2006-01-02 15:04:05"),
		ClassName: className,
		Type:      "INFO",
	}

	l.insert(logMessage)

}

func (l *Log) Error(className, message string) {
	l.currentTime = time.Now()
	logMessage := logDTO{
		Message:   message,
		Timestamp: l.currentTime.Format("2006-01-02 15:04:05"),
		ClassName: className,
		Type:      "ERROR",
	}

	l.insert(logMessage)
}

func (l *Log) getPath() string {
	l.system.Folder("logs")
	return fmt.Sprintf("logs/log_%s.log", l.currentTime.Format("2006-01-02"))
}

func (l *Log) insert(logMessage logDTO) {
	filename := l.getPath()

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(logMessage); err != nil {
		fmt.Printf("Error encoding log message: %v\n", err)
		return
	}
}
