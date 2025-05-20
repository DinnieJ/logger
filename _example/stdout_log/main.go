package main

import (
	"github.com/DinnieJ/logger"
)

func main() {
	l := logger.GetConsoleLogger(&logger.ConsoleLogConfig{
		Name:  "DinnieJ",
		Level: logger.TRACE,
	}, true)
	l.Trace("This is a trace message")
	l.Debug("This is a debug stdout log")
	l.Info("Helloworld")
	l.Warn("This is a warning log")
	l.Error("This is error log")
	l.Fatal("This is a fatal log")
	l.Info(struct {
		ID int
	}{ID: 12})
}
