package base

// LogFormat 日志格式的类型
type LogFormat string

const (
	// FORMAT_TEXT 普通文本日志格式
	FORMAT_TEXY LogFormat = "text"
	// FORMAT_JSON JSON日志格式
	FORMAT_JSON LogFormat = "json"
)

const (
	// TIMESTAMP_FORMAT 时间戳格式化的字符串
	TIMESTAMP_FORMAT = "2006-01-02T15:04:05.999"
)
