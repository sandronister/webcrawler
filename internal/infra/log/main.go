package logCrawller

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type logDTO struct {
	Type      string `json:"type"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	ClassName string `json:"class_name"`
}

type Log struct {
	currentTime time.Time
}

func NewLog() *Log {
	return &Log{}
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
	_, err := os.Stat("logs")
	if os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}
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
