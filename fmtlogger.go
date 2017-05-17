package logger

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	DefaultName = "main"
	Format      = "%s %s %s:%s\n"
	TimeFormat  = "2006-01-02 15:04:05"
)

type qLogger struct {
	wl       int //write level
	subname  string
	strlevel []string //string level
}

func NewFmtLogger(wl string) Logger {
	ql := &qLogger{
		subname:  DefaultName,
		wl:       ERROR,
		strlevel: []string{"FATAL", "ERROR", "INFO", "DEBUG"},
	}
	key := strings.ToUpper(wl)
	for i, v := range ql.strlevel {
		if v == key {
			ql.wl = i
			break
		}
	}
	return ql
}

func (l *qLogger) With(subname string) Logger {
	logger := &qLogger{
		wl:       l.wl,
		subname:  subname,
		strlevel: l.strlevel,
	}
	return logger
}

func (l *qLogger) Debugln(data ...interface{}) {
	l.writeln(DEBUG, data...)
}

func (l *qLogger) Debugf(s string, data ...interface{}) {
	l.writef(DEBUG, s, data...)
}

func (l *qLogger) Infoln(data ...interface{}) {
	l.writeln(INFO, data...)
}

func (l *qLogger) Infof(s string, data ...interface{}) {
	l.writef(INFO, s, data...)
}

func (l *qLogger) Errorln(data ...interface{}) error {
	l.writeln(ERROR, data...)
	return errors.New(fmt.Sprint(data...))
}

func (l *qLogger) Errorf(s string, data ...interface{}) error {
	l.writef(ERROR, s, data...)
	return fmt.Errorf(s, data)
}

func (l *qLogger) Fatalln(data ...interface{}) {
	l.writeln(FATAL, data...)
	os.Exit(1)
}

func (l *qLogger) Fatalf(s string, data ...interface{}) {
	l.writef(FATAL, s, data...)
	os.Exit(1)
}

func (l *qLogger) writeln(level int, data ...interface{}) {
	if l.wl < level {
		return
	}
	fmt.Printf(Format, time.Now().Format(TimeFormat), l.strlevel[level], l.subname, fmt.Sprint(data...))
}

func (l *qLogger) writef(level int, s string, data ...interface{}) {
	if l.wl < level {
		return
	}
	line := fmt.Sprintf(s, data...)
	fmt.Printf(Format, time.Now().Format(TimeFormat), l.strlevel[level], l.subname, line)
}
