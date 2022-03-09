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
	Error(string, ...interface{})
	Warn(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
	log(string, LogLevel, ...interface{})
	Cleanup()
}

// Options is a struct for setting the format of the log messages
type Options struct {
	Format   string
	Position string
}

// Glog type is the logger data for logging concurrently to file
type Glog struct {
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

	var wg sync.WaitGroup

	// set std out as the default glogger
	defaultGlogger = &UnstructureGlog{
		Glog: Glog{
			out:     os.Stderr,
			queue:   make(chan interface{}, 1000),
			wg:      &wg,
			options: []Options{},
			depth:   2,
		},
	}

	wg.Add(1)
	// start a goroutine to write to the file
	// and close the file when the glogger is closed
	// this will block until the queue is empty
	// and all the messages have been written
	// to the file
	// the is the only log writter for glogger
	go func() {
		defer wg.Done()

		for {

			// dequeue new messages to be logged,
			// if chan is closed, break from the loop and close logger
			msg, ok := <-defaultGlogger.queue
			if ok {
				defaultGlogger.out.Write([]byte(msg.(string) + "\n"))
			} else {
				break
			}
		}
	}()
}

// safely close the default logger
func Cleanup() {
	close(defaultGlogger.queue)
	defaultGlogger.wg.Wait()
}

// UnstructureGlogger creates a new Glog object with the given file name and options for formatting
// the log messages. The file is created if it does not exist.
// The file is opened in append mode.
// The log messages are queued and written to the file in a separate goroutine.
// The queue is unbuffered.
func UnstructureGlogger(filePath string, options ...Options) Glogger {

	var wg sync.WaitGroup
	// create the file for logging
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// create a new glogger
	glogger := UnstructureGlog{
		Glog: Glog{
			out:     f,
			queue:   make(chan interface{}, 1000),
			wg:      &wg,
			options: options,
			depth:   1,
		},
	}

	wg.Add(1)
	// start a goroutine to write to the file
	// and close the file when the glogger is closed
	// this will block until the queue is empty
	// and all the messages have been written
	// to the file
	// the is the only log writter for glogger
	go func() {
		defer wg.Done()
		defer f.Close()

		for {

			// dequeue new messages to be logged,
			// if chan is closed, break from the loop and close logger
			msg, ok := <-glogger.queue
			if ok {
				glogger.out.Write([]byte(msg.(string) + "\n"))
			} else {
				break
			}
		}
	}()
	return &glogger
}

// JSONGlogger creates a new Glog object with the given file name and options for formatting
// the log messages. The file is created if it does not exist.
// The file is opened in append mode.
// The log messages are queued and written to the file in a separate goroutine.
// The queue is unbuffered.
func JSONGlogger(filePath string, options ...Options) Glogger {

	var wg sync.WaitGroup
	// create the file for logging
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// create a new glogger
	glogger := JSONGlog{
		Glog: Glog{
			out:     f,
			queue:   make(chan interface{}, 1000),
			wg:      &wg,
			options: options,
			depth:   1,
		},
	}

	wg.Add(1)
	// start a goroutine to write to the file
	// and close the file when the glogger is closed
	// this will block until the queue is empty
	// and all the messages have been written
	// to the file
	// the is the only log writter for glogger
	go func() {
		defer wg.Done()
		defer f.Close()

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
	return &glogger
}

// Error logs the error message to the file
func Error(format string, msg ...interface{}) {
	defaultGlogger.log(format, ERROR, msg...)
}

// Warn logs the warning message to the file
func Warn(format string, msg ...interface{}) {
	defaultGlogger.log(format, WARN, msg...)
}

// Info logs the warning message to the file
func Info(format string, msg ...interface{}) {
	defaultGlogger.log(format, INFO, msg...)
}

// Debug logs the warning message to the file
func Debug(format string, msg ...interface{}) {
	defaultGlogger.log(format, DEBUG, msg...)
}

// returns the string representation of the log level
func (level LogLevel) string() string {
	return string(level)
}
