package logger

const (
	FATAL = iota
	ERROR
	INFO
	DEBUG
)

type Logger interface {
	With(subname string) Logger
	Fatalf(s string, data ...interface{})
	Fatalln(data ...interface{})
	Errorf(s string, data ...interface{}) error
	Errorln(data ...interface{}) error
	Infof(s string, data ...interface{})
	Infoln(data ...interface{})
	Debugf(s string, data ...interface{})
	Debugln(data ...interface{})
}
