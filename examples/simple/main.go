package main

import (
	"github.com/ric-v/glog"
)

// write to stdout
func main() {
	// Cleanup the default concurrent logger
	// this step is required since this is a concurrent logger
	// include this in the main thread, here main func
	defer glog.Cleanup()

	// log the message to the default concurrent logger
	glog.Info("Hello World")
	glog.Warn("Another Hello World")
	glog.Debug("Yet Another Hello World")
	glog.Error("Done for the day")
}

// write to file
// func main() {

// 	logger := glog.NewGlogger("glogger.log")
// 	defer logger.Close()

// 	logger.Log("Hello World")
// 	logger.Log("Another Hello World")
// 	logger.Log("Yet Another Hello World")
// }
