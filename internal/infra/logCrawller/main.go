package logCrawller

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type logDTO struct {
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

func (l *Log) Log(className, message string) {
	l.currentTime = time.Now()
	logMessage := logDTO{
		Message:   message,
		Timestamp: l.currentTime.Format("2006-01-02 15:04:05"),
		ClassName: className,
	}

	l.Insert(logMessage)

}

func (l *Log) Insert(logMessage logDTO) {
	filename := fmt.Sprintf("log_%s.log", l.currentTime.Format("2006-01-02"))

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
