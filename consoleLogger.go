package mylogger

import "fmt"
import "time"
import "runtime"
import "path"

const (
	//DEBUG ...
	DEBUG logLevel = iota
	//TRACE ...
	TRACE
	//INFO ...
	INFO
	//WARNING ...
	WARNING
	//ERROR ...
	ERROR
	//FATAL ...
	FATAL
	//UNKNOWN ...
	UNKNOWN
)

func (c ConsoleLogger) log(status string, format string, s ...interface{}) {
	//var t = time.Now().Format("2006.01.02 15:04:05 UTC")
	var t = time.Now().Format("01-02 15:04:05")
	msg := fmt.Sprintf(format, s...)
	pc, file, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	file = path.Base(file)
	fmt.Printf("[%s] [%s] [%#v|%#v| Line:%d ] %v\n", t, status, funcName, file, line, msg)
}

func (c ConsoleLogger) enable(level logLevel) (ok bool) {
	//fmt.Printf("%d < %d\n", level, int(c.consoleLoggerLevel))
	return level >= c.consoleLoggerLevel
}

//Debug ...
func (c ConsoleLogger) Debug(format string, s ...interface{}) {
	if c.enable(DEBUG) {
		c.log("DEBUG", format, s)
	}
}

//Trace ...
func (c ConsoleLogger) Trace(format string,s ...interface{}) {
	if c.enable(TRACE) {
		c.log("TRACE",format,  s)
	}
}

//Info ...
func (c ConsoleLogger) Info(format string, s ...interface{}) {
	if c.enable(INFO) {
		c.log("INFO",format,  s)
	}
}

//Warning ...
func (c ConsoleLogger) Warning(format string, s ...interface{}) {
	if c.enable(WARNING) {
		c.log("WARNING", format, s)
	}
}

//Error ...
func (c ConsoleLogger) Error(format string, s ...interface{}) {
	if c.enable(ERROR) {
		c.log("ERROR", format, s)
	}
}

//Fatal ...
func (c ConsoleLogger) Fatal(format string, s ...interface{}) {
	if c.enable(FATAL) {
		c.log("FATAL", format, s)
	}
}
