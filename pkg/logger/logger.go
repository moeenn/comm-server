package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	Instance *log.Logger
}

func (l Logger) Info(message string) {
	line := fmt.Sprintf("[INFO] %s", message)
	l.Instance.Println(line)
}

func (l Logger) Warn(message string) {
	line := fmt.Sprintf("[WARN] %s", message)
	l.Instance.Println(line)
}

func (l Logger) Error(message string) {
	line := fmt.Sprintf("[ERROR] %s", message)
	l.Instance.Println(line)
}

func New() *Logger {
	return &Logger{
		Instance: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}
}
