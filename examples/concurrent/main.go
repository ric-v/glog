package main

import (
	"strconv"
	"sync"
	"time"

	"github.com/ric-v/glog"
)

func main() {

	logger := glog.NewGlogger("glogger.log")

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()

			logger.Log("Hello World " + strconv.Itoa(i))
			logger.Log("Another Hello World " + strconv.Itoa(i))
			logger.Log("Yet Another Hello World " + strconv.Itoa(i))
		}(i, &wg)
	}
	wg.Wait()
	time.Sleep(time.Millisecond)
}
