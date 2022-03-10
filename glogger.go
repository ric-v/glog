package glog

import (
	"io"
	"os"
	"sync"
)

const (
	ERROR LogLevel = "ERROR"
	INFO  LogLevel = "INFO"
	DEBUG LogLevel = "DEBUG"
	WARN  LogLevel = "WARN"
)

type LogLevel string

// Glogger is the interface for controlling logging
type Glogger interface {
	Error(...interface{})         // error logger without format
	Warn(...interface{})          // warning logger without format
	Info(...interface{})          // info logger without format
	Debug(...interface{})         // debug logger without format
	log(LogLevel, ...interface{}) // logger without format

	Errorf(string, ...interface{})         // error logger with format
	Warnf(string, ...interface{})          // warning logger with format
	Infof(string, ...interface{})          // info logger with format
	Debugf(string, ...interface{})         // debug logger with format
	logf(string, LogLevel, ...interface{}) // logger with format

	Cleanup() // safely close the glogger
}

// Options is a struct for setting the format of the log messages
type Options struct {
	Format   string
	Position string
}

// glog type is the logger data for logging concurrently to file
type glog struct {
	out     io.Writer        // writer for logging to file / cmd line
	queue   chan interface{} // queue for logging
	wg      *sync.WaitGroup  // wait group for logging
	options []Options        // options for formatting the log messages
	depth   int              // depth of the function caller
}

// defaultGlogger is the default glogger
var defaultGlogger *UnstructureGlog

// initialize a default logger when user doesnot want a custom logger
func init() {
	NewDefaultGlogger()
}

// NewDefaultGlogger creates a new Glog object for stdout and options for formatting
func NewDefaultGlogger() Glogger {

	var wg sync.WaitGroup

	// set std out as the default glogger
	defaultGlogger = &UnstructureGlog{
		glog: glog{
			out:     os.Stderr,
			queue:   make(chan interface{}, 1000),
			wg:      &wg,
			options: []Options{},
			depth:   2,
		},
	}

	// start a goroutine to write to the file
	sink(&defaultGlogger.glog)
	return defaultGlogger
}

// safely close the default logger
func Cleanup() {
	close(defaultGlogger.queue)
	defaultGlogger.wg.Wait()
}

// NewUnstructureGlogger creates a new Glog object with the given file name and options for formatting
// the log messages. The file is created if it does not exist.
// The file is opened in append mode.
// The log messages are queued and written to the file in a separate goroutine.
// The queue is unbuffered.
func NewUnstructureGlogger(filePath string, options ...Options) Glogger {

	// create new unstructured glogger
	glogger, err := newUnstructuredGlogger(filePath, options...)
	if err != nil {
		return nil
	}

	// start a goroutine to write to the file
	sink(&glogger.glog)
	return glogger
}

// NewJSONGlogger creates a new Glog object with the given file name and options for formatting
// the log messages. The file is created if it does not exist.
// The file is opened in append mode.
// The log messages are queued and written to the file in a separate goroutine.
// The queue is unbuffered.
func NewJSONGlogger(filePath string, options ...Options) Glogger {

	// create new unstructured glogger
	glogger, err := newJSONGlogger(filePath, options...)
	if err != nil {
		return nil
	}

	// start a goroutine to write to the file
	sink(&glogger.glog)
	return glogger
}

// sink for logging asynchronously
func sink(glogger *glog) {

	glogger.wg.Add(1)
	// start a goroutine to write to the file
	// and close the file when the glogger is closed
	// this will block until the queue is empty
	// and all the messages have been written
	// to the file
	// the is the only log writter for glogger
	go func() {
		defer glogger.wg.Done()
		// defer glogger.out.Close()

		for {

			// dequeue new messages to be logged,
			// if chan is closed, break from the loop and close logger
			msg, ok := <-glogger.queue
			if ok {
				glogger.out.Write(msg.([]byte))
			} else {
				break
			}
		}
	}()
}

// Error logs the error message to the file
func Error(msg ...interface{}) {
	defaultGlogger.log(ERROR, msg...)
}

// Warn logs the warning message to the file
func Warn(msg ...interface{}) {
	defaultGlogger.log(WARN, msg...)
}

// Info logs the warning message to the file
func Info(msg ...interface{}) {
	defaultGlogger.log(INFO, msg...)
}

// Debug logs the warning message to the file
func Debug(msg ...interface{}) {
	defaultGlogger.log(DEBUG, msg...)
}

// Error logs the error message to the file
func Errorf(format string, msg ...interface{}) {
	defaultGlogger.logf(format, ERROR, msg...)
}

// Warn logs the warning message to the file
func Warnf(format string, msg ...interface{}) {
	defaultGlogger.logf(format, WARN, msg...)
}

// Info logs the warning message to the file
func Infof(format string, msg ...interface{}) {
	defaultGlogger.logf(format, INFO, msg...)
}

// Debug logs the warning message to the file
func Debugf(format string, msg ...interface{}) {
	defaultGlogger.logf(format, DEBUG, msg...)
}

// returns the string representation of the log level
func (level LogLevel) string() string {
	return string(level)
}
