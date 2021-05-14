package base

import "github.com/rszhh/gowcer/log/field"

// Option 日志记录器的选项
type Option interface {
	Name() string
}

// MyLogger 日志记录器
type MyLogger interface {
	Name() string
	Level() LogLevel
	Format() LogFormat
	Options() []Option

	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Debugln(v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Infoln(v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Warnln(v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})

	// WithFields 会增加需记录的额外字段。
	WithFields(fields ...field.Field) MyLogger
}
