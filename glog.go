package glog

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func (g *Glog) Log(msg string) {
	_, filename, line, _ := runtime.Caller(1)
	g.queue <- time.Now().Format("[2006-01-02 15:04:05.000 MST]") +
		"\t(" + strings.Join(strings.Split(filename, "/")[len(strings.Split(filename, "/"))-2:], "/") +
		": " + fmt.Sprint(line) + ")\t" + msg
}

func (g *Glog) Close() {
	close(g.queue)
}
