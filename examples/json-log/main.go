package main

import "github.com/ric-v/glog"

func main() {

	logger := glog.NewJSONGlogger("glogger.json")
	defer logger.Cleanup()

	logger.Info("", "Hello", "World")
	logger.Warn("", "Another", "Hello", "World")
	logger.Error("", "Error", "Hello", "World")
	logger.Debug("", "Debug", "Hello", "World")

}
