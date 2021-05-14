package base

// LogLevel 日志输出级别
type LogLevel uint8

const (
	// LEVEL_DEBUG 调试级别，是最低的调试等级
	LEVEL_DEBUG LogLevel = iota + 1
	// LEVEL_INFO 信息级别，是最常用的日志等级
	LEVEL_INFO
	// LEVEL_WARN 警告级别，是适合输出到错误输出的日志等级
	LEVEL_WARN
	// LEVEL_ERROR 普通错误级别，输出到错误输出的日志等级
	LEVEL_ERROR
	// LEVEL_FATAL 致命错误级别，输出到错误输出的日志等级
	// 此种错误级别一旦出现就意味着 os.Exit(1) 立即被调用
	LEVEL_FATAL
	// LEVEL_PANIC 恐慌级别，是最高的日志等级
	LEVEL_PANIC
)
