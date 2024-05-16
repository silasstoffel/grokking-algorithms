package logger

import "log"

type ILogger interface {
	Info(message string, data interface{})
}

type Logger struct {
	format string
}

func NewLogger() *Logger {
	return &Logger{
		format: "text",
	}
}

func (*Logger) Info(message string, data interface{}) {
	log.Println(message, data)
}
