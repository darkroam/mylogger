package mylogger

import "strings"

//MyLogger ...
type MyLogger interface {
	Debug(s interface{})
	Trace(s interface{})
	Info(s interface{})
	Warning(s interface{})
	Error(s interface{})
	Fatal(s interface{})
}

type logLevel uint16

//ConsoleLogger ...
type ConsoleLogger struct {
	consoleLoggerLevel logLevel
}

// NewConsoleLogger ...
func NewConsoleLogger(level string) (c *ConsoleLogger) {
	c = new(ConsoleLogger)
	level = strings.ToUpper(level)
	//fmt.Printf("level = %s\n", level)
	switch level {
	case "DEBUG":
		c.consoleLoggerLevel = DEBUG
	case "TRACE":
		c.consoleLoggerLevel = TRACE
	case "INFO":
		c.consoleLoggerLevel = INFO
	case "WARNING":
		c.consoleLoggerLevel = WARNING
	case "ERROR":
		c.consoleLoggerLevel = ERROR
	case "FATAL":
		c.consoleLoggerLevel = FATAL
	default:
		c.consoleLoggerLevel = UNKNOWN
	}
	return
}
