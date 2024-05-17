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
	if data != nil {
		log.Println(message, data)
	} else {
		log.Println(message)
	}
}
