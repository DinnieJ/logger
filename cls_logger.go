package logger

import (
	"log"
	"os"
)

var (
	logNameColor = "\x1b[35m\x1b[1m"
	resetColor   = "\x1b[0m"
	flagColor    = "\x1b[234m\x1b[1m"
	traceColor   = "\x1b[36m\x1b[1m"
	debugColor   = "\x1b[34m\x1b[4m\x1b[1m"
	infoColor    = "\x1b[32m\x1b[1m"
	warnColor    = "\x1b[33m\x1b[1m"
	errorColor   = "\x1b[31m\x1b[1m"
	fatalColor   = "\x1b[41m\x1b[1m"
)

var LogLevelStrMap map[LogLevel]string = map[LogLevel]string{
	DEBUG: debugColor + "[" + LOG_LEVEL_DEBUG + "]" + resetColor,
	TRACE: traceColor + "[" + LOG_LEVEL_TRACE + "]" + resetColor,
	INFO:  infoColor + "[" + LOG_LEVEL_INFO + "]" + resetColor,
	WARN:  warnColor + "[" + LOG_LEVEL_WARN + "]" + resetColor,
	ERROR: errorColor + "[" + LOG_LEVEL_ERROR + "]" + resetColor,
	FATAL: fatalColor + "[" + LOG_LEVEL_FATAL + "]" + resetColor,
}

type ConsoleLogConfig struct {
	Name  string
	Level LogLevel
	Color bool
}

func (c *ConsoleLogConfig) GetLogName() string {
	return c.Name
}

func (c *ConsoleLogConfig) GetLogLevel() LogLevel {
	return c.Level
}

type ClsLogger struct {
	config ConsoleLogConfig
	logger *log.Logger
}

func NewConsoleLogger(config LogConfig, color bool) *ClsLogger {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	return &ClsLogger{
		config: ConsoleLogConfig{
			Name:  config.GetLogName(),
			Level: config.GetLogLevel(),
			Color: color,
		},
		logger: logger,
	}
}

func (c *ClsLogger) log(msg any, level LogLevel) error {
	if c.config.Level > level {
		return nil
	}
	c.logger.SetPrefix(flagColor)
	logNameColor := logNameColor + "[" + c.config.GetLogName() + "]" + resetColor
	c.logger.Printf("%s%s %s\t|%s", resetColor, logNameColor, LogLevelStrMap[level], msg)

	return nil
}

func (c *ClsLogger) Debug(msg any) error {
	return c.log(msg, DEBUG)
}

func (c *ClsLogger) Trace(msg any) error {
	return c.log(msg, TRACE)
}

func (c *ClsLogger) Info(msg any) error {
	return c.log(msg, INFO)
}

func (c *ClsLogger) Warn(msg any) error {
	return c.log(msg, WARN)
}

func (c *ClsLogger) Error(msg any) error {
	return c.log(msg, ERROR)
}

func (c *ClsLogger) Fatal(msg any) error {
	return c.log(msg, FATAL)
}
