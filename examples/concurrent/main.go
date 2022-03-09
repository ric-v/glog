package main

import (
	"strconv"
	"sync"

	"github.com/ric-v/glog"
)

func main() {

	logger := glog.UnstructureGlogger("glogger.log")
	defer logger.Close()

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			logger.Info("Hello World " + strconv.Itoa(i))
		}(i, &wg)
	}
	wg.Wait()
}
