package log

import (
	"io"
	"os"
	"sync"

	"github.com/rszhh/gowcer/log/base"
	"github.com/rszhh/gowcer/log/logrus"
)

// rwm 日志记录器创建器映射的专用锁
var rwm sync.RWMutex

// LoggerCreator 日志记录器的创建器
type LoggerCreator func(
	level base.LogLevel,
	format base.LogFormat,
	writer io.Writer,
	options []base.Option) base.MyLogger

// loggerCreatorMap 日志记录器创建器的映射
var loggerCreatorMap = map[base.LogType]LoggerCreator{}

// RegisterLogger 注册日志记录器
// 似乎没用到

// Logger 新建一个日志记录器
func Logger(
	logType base.LogType,
	level base.LogLevel,
	format base.LogFormat,
	writer io.Writer,
	options []base.Option) base.MyLogger {
	var logger base.MyLogger
	rwm.RLock()
	creater, ok := loggerCreatorMap[logType]
	rwm.RUnlock()
	// fmt.Println(creater)
	if ok {
		// fmt.Println("1")
		logger = creater(level, format, writer, options)
	} else {
		// fmt.Println("2")
		logger = logrus.NewLoggerBy(level, format, writer, options)
	}
	return logger
}

// DLogger 返回一个新的默认日志记录器
func DLogger() base.MyLogger {
	return Logger(
		base.TYPE_LOGRUS,
		base.LEVEL_INFO,
		base.FORMAT_TEXY,
		os.Stdout,
		nil)
}
