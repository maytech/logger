package main

import (
	"fmt"
	"io/ioutil"
	syslog "log"
	"os"
	"strings"
)

type Level int

const (
	Fatal Level = iota
	Error
	Info
	Debug
)

func (l Level) String() string {
	switch l {
	case Fatal:
		return "fatal"
	case Error:
		return "error"
	case Info:
		return "info"
	case Debug:
		return "debug"
	}
	return ""
}

func LevelFromString(l string) Level {
	l = strings.ToLower(l)
	switch l {
	case "fatal":
		return Fatal
	case "error":
		return Error
	case "info":
		return Info
	case "debug":
		return Debug
	}
	return Info
}

type Logger struct {
	logger *syslog.Logger
	level  Level
}

func NewFileLogger(path, prefix string, level Level) (*Logger, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	l := syslog.New(f, prefix, syslog.LstdFlags)
	return &Logger{l, level}, nil
}

func NewStdoutLogger(prefix string, level Level) (*Logger, error) {
	l := syslog.New(os.Stdout, prefix, syslog.LstdFlags)
	return &Logger{l, level}, nil
}

func NewEmptyLogger(prefix string, level Level) *Logger {
	l := syslog.New(ioutil.Discard, prefix, syslog.LstdFlags)
	return &Logger{l, level}
}

func (l *Logger) Fatalf(s string, data ...interface{}) {
	if l.level >= Fatal {
		l.logf(s, Fatal, data...)
		os.Exit(1)
	}
}

func (l *Logger) Fatalln(data ...interface{}) {
	if l.level >= Fatal {
		l.logln(Fatal, data...)
		os.Exit(1)
	}
}

func (l *Logger) Fatal(data ...interface{}) {
	l.Fatalln(data...)
}

func (l *Logger) Errorf(s string, data ...interface{}) {
	if l.level >= Error {
		l.logf(s, Error, data...)
	}
}

func (l *Logger) Errorln(data ...interface{}) {
	if l.level >= Error {
		l.logln(Error, data...)
	}
}

func (l *Logger) Error(data ...interface{}) {
	l.Errorln(data...)
}

func (l *Logger) Infof(s string, data ...interface{}) {
	if l.level >= Info {
		l.logf(s, Info, data...)
	}
}

func (l *Logger) Infoln(data ...interface{}) {
	if l.level >= Info {
		l.logln(Info, data...)
	}
}

func (l *Logger) Info(data ...interface{}) {
	l.Infoln(data...)
}

func (l *Logger) Debugf(s string, data ...interface{}) {
	if l.level >= Debug {
		l.logf(s, Debug, data...)
	}
}

func (l *Logger) Debugln(data ...interface{}) {
	if l.level >= Debug {
		l.logln(Debug, data...)
	}
}

func (l *Logger) Debug(data ...interface{}) {
	l.Debugln(data...)
}

func (l *Logger) logf(s string, level Level, data ...interface{}) {
	l.logger.Printf("[%s] %s\n", level, fmt.Sprintf(s, data...))
}

func (l *Logger) logln(level Level, data ...interface{}) {
	l.logger.Printf("[%s] %s\n", level, fmt.Sprint(data...))
}
