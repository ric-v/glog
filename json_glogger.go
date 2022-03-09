package glog

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// JSONGlog  type is the logger data for logging concurrently to file
// easyjson:skip
type JSONGlog struct {
	Glog
}

// easyjson:json
type LoggerJson struct {
	Time string                 `json:"time"`
	Type string                 `json:"type"`
	File string                 `json:"file"`
	Line string                 `json:"line"`
	Msg  map[string]interface{} `json:"msg"`
}

// Error logs the error message to the file
func (g *JSONGlog) Error(format string, msg ...interface{}) {
	g.log(format, ERROR, msg...)
}

// Warn logs the warning message to the file
func (g *JSONGlog) Warn(format string, msg ...interface{}) {
	g.log(format, WARN, msg...)
}

// Info logs the warning message to the file
func (g *JSONGlog) Info(format string, msg ...interface{}) {
	g.log(format, INFO, msg...)
}

// Debug logs the warning message to the file
func (g *JSONGlog) Debug(format string, msg ...interface{}) {
	g.log(format, DEBUG, msg...)
}

// log runs the custom logger and logs the message to the file
func (g *JSONGlog) log(format string, level LogLevel, msg ...interface{}) {

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
