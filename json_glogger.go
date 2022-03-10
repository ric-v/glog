package glog

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

// JSONGlog  type is the logger data for logging concurrently to file
// easyjson:skip
type JSONGlog struct {
	glog
}

// easyjson:json
type LoggerJson struct {
	Time string                 `json:"time"`
	Type string                 `json:"type"`
	File string                 `json:"file"`
	Line string                 `json:"line"`
	Msg  map[string]interface{} `json:"msg"`
}

// newJSONGlogger returns a new JSON glogger
func newJSONGlogger(filePath string, options ...Options) (*JSONGlog, error) {

	var wg sync.WaitGroup
	// create the file for logging
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// create a new glogger
	return &JSONGlog{
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
func (g *JSONGlog) Error(msg ...interface{}) {
	g.log(ERROR, msg...)
}

// Warn logs the warning message to the file
func (g *JSONGlog) Warn(msg ...interface{}) {
	g.log(WARN, msg...)
}

// Info logs the warning message to the file
func (g *JSONGlog) Info(msg ...interface{}) {
	g.log(INFO, msg...)
}

// Debug logs the warning message to the file
func (g *JSONGlog) Debug(msg ...interface{}) {
	g.log(DEBUG, msg...)
}

// log runs the custom logger and logs the message to the file
func (g *JSONGlog) log(level LogLevel, msg ...interface{}) {
	g.jsonLogger(level, msg...)
}

// Error logs the error message to the file
func (g *JSONGlog) Errorf(format string, msg ...interface{}) {
	g.logf(format, ERROR, msg...)
}

// Warn logs the warning message to the file
func (g *JSONGlog) Warnf(format string, msg ...interface{}) {
	g.logf(format, WARN, msg...)
}

// Info logs the warning message to the file
func (g *JSONGlog) Infof(format string, msg ...interface{}) {
	g.logf(format, INFO, msg...)
}

// Debug logs the warning message to the file
func (g *JSONGlog) Debugf(format string, msg ...interface{}) {
	g.logf(format, DEBUG, msg...)
}

// log runs the custom logger and logs the message to the file
func (g *JSONGlog) logf(format string, level LogLevel, msg ...interface{}) {
	g.jsonLogger(level, msg...)
}

func (g *JSONGlog) jsonLogger(level LogLevel, msg ...interface{}) {

	// get the function caller info
	_, filename, line, _ := runtime.Caller(g.depth)

	// generate json data for logging
	log, err := LoggerJson{
		Time: time.Now().Format(time.RFC3339),
		Type: level.string(),
		File: strings.Join(strings.Split(filename, "/")[len(strings.Split(filename, "/"))-2:], "/"),
		Line: fmt.Sprintf("%d", line),
		Msg: func() map[string]interface{} {
			var logData = make(map[string]interface{})
			for i := 0; i < len(msg)-1; i += 2 {
				logData[fmt.Sprint(msg[i])] = msg[i+1]
			}
			if len(msg)%2 != 0 {
				logData[fmt.Sprint(msg[len(msg)-1])] = ""
			}
			return logData
		}(),
	}.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}

	// add new log message to the queue
	g.queue <- append(log, '\n')
}

// safely close the custom logger
func (g *JSONGlog) Cleanup() {
	close(g.queue)
	g.wg.Wait()
}
