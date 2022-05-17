package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type LoggerService interface {
	SetTimeLoyaut(t string) LoggerService
	SetServiceName(name string) LoggerService
	SetLevel(level int) LoggerService
	SetOutput(w io.Writer) LoggerService

	Debug(message string)
	Debugf(format string, values ...interface{})
	Info(message string)
	Infof(format string, values ...interface{})
	Warning(message string)
	Warningf(format string, values ...interface{})
	Error(message string)
	Errorf(format string, values ...interface{})
	Fatal(message string)
	Fatalf(format string, values ...interface{})
}

type service struct {
	timeLoyaut  string
	serviceName string
	level       int
	out         io.Writer
	eventChan   chan *event
}

type event struct {
	message string
	time    time.Time
	level   int
}

const (
	DebugLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

var (
	DefaultTimeLoyaut  = time.RFC1123
	DefaultServiceName = "logger"
	DefaultOutType     = os.Stdout
	DefaultLevel       = ErrorLevel

	levelMap = map[int]string{
		DebugLevel:   "DEBUG",
		InfoLevel:    "INFO",
		WarningLevel: "WARNING",
		ErrorLevel:   "ERROR",
		FatalLevel:   "FATAL",
	}
)

func NewLoggerService() LoggerService {
	s := &service{
		timeLoyaut:  DefaultTimeLoyaut,
		serviceName: DefaultServiceName,
		out:         DefaultOutType,
		eventChan:   make(chan *event),
	}
	defer func() {
		go s.run()
	}()
	return s
}

func (s *service) SetTimeLoyaut(t string) LoggerService {
	s.timeLoyaut = t
	return s
}

func (s *service) SetServiceName(name string) LoggerService {
	s.serviceName = name
	return s
}

func (s *service) SetOutput(w io.Writer) LoggerService {
	s.out = w
	return s
}

func (s *service) SetLevel(level int) LoggerService {
	if _, ok := levelMap[level]; !ok {
		level = ErrorLevel
	}
	s.level = level
	return s
}

func (s *service) Debug(message string) {
	s.eventChan <- &event{message, time.Now(), DebugLevel}
}

func (s *service) Debugf(format string, values ...interface{}) {
	s.eventChan <- &event{fmt.Sprintf(format, values...), time.Now(), DebugLevel}
}

func (s *service) Info(message string) {
	s.eventChan <- &event{message, time.Now(), InfoLevel}
}

func (s *service) Infof(format string, values ...interface{}) {
	s.eventChan <- &event{fmt.Sprintf(format, values...), time.Now(), InfoLevel}
}

func (s *service) Warning(message string) {
	s.eventChan <- &event{message, time.Now(), WarningLevel}
}

func (s *service) Warningf(format string, values ...interface{}) {
	s.eventChan <- &event{fmt.Sprintf(format, values...), time.Now(), WarningLevel}
}

func (s *service) Error(message string) {
	s.eventChan <- &event{message, time.Now(), ErrorLevel}
}

func (s *service) Errorf(format string, values ...interface{}) {
	s.eventChan <- &event{fmt.Sprintf(format, values...), time.Now(), ErrorLevel}
}

func (s *service) Fatal(message string) {
	s.print(&event{message, time.Now(), FatalLevel})
	os.Exit(1)
}

func (s *service) Fatalf(format string, values ...interface{}) {
	s.print(&event{fmt.Sprintf(format, values...), time.Now(), FatalLevel})
	os.Exit(1)
}

func (s *service) run() {
	for {
		e, ok := <-s.eventChan
		if !ok {
			panic("logger channel is closed")
		}
		if e.level < s.level {
			continue
		}
		s.print(e)
	}
}

func (s *service) print(e *event) {
	if _, err := fmt.Fprintf(s.out, "%s | %s | %s | %s\n", s.serviceName, e.time.Format(s.timeLoyaut), levelMap[e.level], e.message); err != nil {
		panic(err)
	}
}
