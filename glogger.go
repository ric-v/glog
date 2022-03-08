package glog

import (
	"os"
)

type Glogger interface {
	Log(msg string)
	Close()
}

// Glog type is the logger data for logging concurrently to file
type Glog struct {
	file  *os.File
	queue chan interface{}
}

type Options struct {
	Format   string
	Position string
}

// NewGlogger creates a new Glog object with the given file name and options for formatting
// the log messages. The file is created if it does not exist.
// The file is opened in append mode.
// The log messages are queued and written to the file in a separate goroutine.
// The queue is unbuffered.
func NewGlogger(filePath string, options ...Options) Glogger {

	// create the file for logging
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// create a new glogger
	glogger := Glog{
		file:  f,
		queue: make(chan interface{}, 100),
	}

	// start a goroutine to write to the file
	// and close the file when the glogger is closed
	// this will block until the queue is empty
	// and all the messages have been written
	// to the file
	// the is the only log writter for glogger
	go func() {
		defer f.Close()

		for {
			msg := <-glogger.queue
			f.Write([]byte(msg.(string) + "\n"))
		}
	}()

	// return the glogger
	return &glogger
}
