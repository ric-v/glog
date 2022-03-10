package glog

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

// UnstructureGlog type is the logger data for logging concurrently to file
type UnstructureGlog struct {
	glog
}

// newUnstructuredGlogger returns a new unstructured glogger
func newUnstructuredGlogger(filePath string, options ...Options) (*UnstructureGlog, error) {

	var wg sync.WaitGroup
	// create the file for logging
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// create a new glogger
	return &UnstructureGlog{
		glog: glog{
			out:     f,
			queue:   make(chan interface{}, 1000),
			wg:      &wg,
			options: options,
			depth:   2,
		},
	}, nil
}

// Error logs the error message to the file
func (g *UnstructureGlog) Error(msg ...interface{}) {
	g.log(ERROR, msg...)
}

// Warn logs the warning message to the file
func (g *UnstructureGlog) Warn(msg ...interface{}) {
	g.log(WARN, msg...)
}

// Info logs the warning message to the file
func (g *UnstructureGlog) Info(msg ...interface{}) {
	g.log(INFO, msg...)
}

// Debug logs the warning message to the file
func (g *UnstructureGlog) Debug(msg ...interface{}) {
	g.log(DEBUG, msg...)
}

// log runs the custom logger and logs the message to the file
func (g *UnstructureGlog) log(level LogLevel, msg ...interface{}) {
	g.unstructuredLogger(level, fmt.Sprint(msg...))
}

// Error logs the error message to the file
func (g *UnstructureGlog) Errorf(format string, msg ...interface{}) {
	g.logf(format, ERROR, msg...)
}

// Warn logs the warning message to the file
func (g *UnstructureGlog) Warnf(format string, msg ...interface{}) {
	g.logf(format, WARN, msg...)
}

// Info logs the warning message to the file
func (g *UnstructureGlog) Infof(format string, msg ...interface{}) {
	g.logf(format, INFO, msg...)
}

// Debug logs the warning message to the file
func (g *UnstructureGlog) Debugf(format string, msg ...interface{}) {
	g.logf(format, DEBUG, msg...)
}

// log runs the custom logger and logs the message to the file
func (g *UnstructureGlog) logf(format string, level LogLevel, msg ...interface{}) {
	g.unstructuredLogger(level, fmt.Sprintf(format, msg...))
}

func (g *UnstructureGlog) unstructuredLogger(level LogLevel, msg string) {

	// get the function caller info
	_, filename, line, _ := runtime.Caller(g.depth)

	// add new log message to the queue
	g.queue <- []byte(time.Now().Format("[2006-01-02 15:04:05.000 MST]") +
		"\t" + level.string() + "\t(" + strings.Join(strings.Split(filename, "/")[len(strings.Split(filename, "/"))-2:], "/") +
		": " + fmt.Sprint(line) + ")\t" + msg)
}

// safely close the custom logger
func (g *UnstructureGlog) Cleanup() {
	close(g.queue)
	g.wg.Wait()
}
