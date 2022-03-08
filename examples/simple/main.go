package main

import (
	"time"

	"github.com/ric-v/glog"
)

func main() {

	logger := glog.NewGlogger("glogger.log")

	logger.Log("Hello World")
	logger.Log("Another Hello World")
	logger.Log("Yet Another Hello World")
	time.Sleep(time.Millisecond)
}
