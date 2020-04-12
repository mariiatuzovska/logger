package main

import (
	"fmt"
	"log"
	"os"
)

type Logger struct{}

func New() *Logger {
	return new(Logger)
}

func (l *Logger) Trace(format string, v ...interface{}) {
	log.Println("| TRACE | ", format + "\n", v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	log.Println(" | INFO | ", format + "\n", v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	log.Println(" | WARNING | ", format + "\n", v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	log.Println(" | ERROR | ", format + "\n", v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	log.Println(" | FATAL | ", format + "\n", v...)
	os.Exit(1)
}