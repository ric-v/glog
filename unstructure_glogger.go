package glog

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// UnstructureGlog type is the logger data for logging concurrently to file
type UnstructureGlog struct {
	Glog
}

// Error logs the error message to the file
func (g *UnstructureGlog) Error(format string, msg ...interface{}) {
	g.log(format, ERROR, msg...)
}

// Warn logs the warning message to the file
func (g *UnstructureGlog) Warn(format string, msg ...interface{}) {
	g.log(format, WARN, msg...)
}

// Info logs the warning message to the file
func (g *UnstructureGlog) Info(format string, msg ...interface{}) {
	g.log(format, INFO, msg...)
}

// Debug logs the warning message to the file
func (g *UnstructureGlog) Debug(format string, msg ...interface{}) {
	g.log(format, DEBUG, msg...)
}

// log runs the custom logger and logs the message to the file
func (g *UnstructureGlog) log(format string, level LogLevel, msg ...interface{}) {

	// get the function caller info
	_, filename, line, _ := runtime.Caller(g.depth)

	// add new log message to the queue
	g.queue <- time.Now().Format("[2006-01-02 15:04:05.000 MST]") +
		"\t" + level.string() + "\t(" + strings.Join(strings.Split(filename, "/")[len(strings.Split(filename, "/"))-2:], "/") +
		": " + fmt.Sprint(line) + ")\t" + fmt.Sprintf(format, msg...)
}

// safely close the custom logger
func (g *UnstructureGlog) Cleanup() {
	close(g.queue)
	g.wg.Wait()
}
