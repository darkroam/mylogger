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

func (c ConsoleLogger) log(status string, s interface{}) {
	//var t = time.Now().Format("2006.01.02 15:04:05 UTC")
	var t = time.Now().Format("01-02 15:04:05")
	pc, file, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	file = path.Base(file)
	fmt.Printf("[%s] [%s] [%#v|%#v| Line:%d ] %v\n", t, status, funcName, file, line, s)
}

func (c ConsoleLogger) enable(level logLevel) (ok bool) {
	//fmt.Printf("%d < %d\n", level, int(c.consoleLoggerLevel))
	return level >= c.consoleLoggerLevel
}

//Debug ...
func (c ConsoleLogger) Debug(s interface{}) {
	if c.enable(DEBUG) {
		c.log("DEBUG", s)
	}
}

//Trace ...
func (c ConsoleLogger) Trace(s interface{}) {
	if c.enable(TRACE) {
		c.log("TRACE", s)
	}
}

//Info ...
func (c ConsoleLogger) Info(s interface{}) {
	if c.enable(INFO) {
		c.log("INFO", s)
	}
}

//Warning ...
func (c ConsoleLogger) Warning(s interface{}) {
	if c.enable(WARNING) {
		c.log("WARNING", s)
	}
}

//Error ...
func (c ConsoleLogger) Error(s interface{}) {
	if c.enable(ERROR) {
		c.log("ERROR", s)
	}
}

//Fatal ...
func (c ConsoleLogger) Fatal(s interface{}) {
	if c.enable(FATAL) {
		c.log("FATAL", s)
	}
}
