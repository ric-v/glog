package main

import "github.com/ric-v/glog"

func main() {

	logger := glog.JSONGlogger("glogger.json")
	defer logger.Close()

	logger.Info("", "Hello", "World")
	logger.Warn("", "Another", "Hello", "World")
	logger.Error("", "Error", "Hello", "World")
	logger.Debug("", "Debug", "Hello", "World")

}
