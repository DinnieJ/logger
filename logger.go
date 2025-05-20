package logger

type Logger interface {
	Trace(msg any) error
	Debug(msg any) error
	Info(msg any) error
	Warn(msg any) error
	Error(msg any) error
	Fatal(msg any) error
}
type LogLevel uint

const (
	TRACE LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type LogConfig interface {
	GetLogName() string
	GetLogLevel() LogLevel
}

const LOG_LEVEL_TRACE = "TRACE"
const LOG_LEVEL_DEBUG = "DEBUG"
const LOG_LEVEL_INFO = "INFO"
const LOG_LEVEL_WARN = "WARNING"
const LOG_LEVEL_ERROR = "ERROR"
const LOG_LEVEL_FATAL = "FATAL"

func GetConsoleLogger(cfg LogConfig, color bool) Logger {
	return NewConsoleLogger(cfg, color)
}
