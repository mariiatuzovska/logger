package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct{}

func New() *Logger {
	return new(Logger)
}

func (l *Logger) Trace(format string, values ...interface{}) {
	log.Println(message("TRACE", format, values))
}

func (l *Logger) Info(format string, values ...interface{}) {
	log.Println(message("INFO", format, values))
}

func (l *Logger) Warn(format string, values ...interface{}) {
	log.Println(message("WARNING", format, values))
}

func (l *Logger) Error(format string, values ...interface{}) {
	log.Println(message("ERROR", format, values))
}

func (l *Logger) Fatal(format string, values ...interface{}) {
	log.Println(message("FATAL", format, values))
	os.Exit(1)
}

func message(TYPE, FORMAT string, values ...interface{}) string {
	return fmt.Sprintf(" | %s | %s", TYPE, FORMAT, values)
}
