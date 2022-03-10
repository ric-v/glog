package glog

import (
	"os"
	"sync"
	"testing"
)

func TestUnstructureGlog_log(t *testing.T) {
	type args struct {
		format string
		level  LogLevel
		msg    []interface{}
	}
	tests := []struct {
		name string
		g    *UnstructureGlog
		args args
	}{
		{
			name: "success",
			g: &UnstructureGlog{
				glog: glog{
					out:   os.Stdout,
					queue: make(chan interface{}, 100),
					wg:    &sync.WaitGroup{},
				},
			},
			args: args{
				format: "%s",
				level:  ERROR,
				msg:    []interface{}{"Success"},
			},
		},
	}
	for _, tt := range tests {
		defer tt.g.Cleanup()
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Error(tt.args.msg...)
			tt.g.Info(tt.args.msg...)
			tt.g.Warn(tt.args.msg...)
			tt.g.Debug(tt.args.msg...)
			tt.g.log(tt.args.level, tt.args.msg...)
		})
	}
	os.RemoveAll(testFile)
}

func TestUnstructureGlog_logf(t *testing.T) {
	type args struct {
		format string
		level  LogLevel
		msg    []interface{}
	}
	tests := []struct {
		name string
		g    *UnstructureGlog
		args args
	}{
		{
			name: "success",
			g: &UnstructureGlog{
				glog: glog{
					out:   os.Stdout,
					queue: make(chan interface{}, 100),
					wg:    &sync.WaitGroup{},
				},
			},
			args: args{
				format: "%s",
				level:  ERROR,
				msg:    []interface{}{"Success"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.g.Errorf(tt.args.format, tt.args.msg...)
			tt.g.Infof(tt.args.format, tt.args.msg...)
			tt.g.Warnf(tt.args.format, tt.args.msg...)
			tt.g.Debugf(tt.args.format, tt.args.msg...)
			tt.g.logf(tt.args.format, tt.args.level, tt.args.msg...)
		})
	}
	os.RemoveAll(testFile)
}

// benchmarking
func BenchmarkUnstructureGlog_logf(b *testing.B) {
	g := NewUnstructureGlogger(testFile)
	for i := 0; i < b.N; i++ {
		g.logf("%s, %v, %d, %t, %s, %v, %d, %t", ERROR, "Success", g, 1, true, "Success", g, 1, true)
	}
	g.Cleanup()
	os.RemoveAll(testFile)
}

// benchmarking
func BenchmarkUnstructureGlog_log(b *testing.B) {
	g := NewUnstructureGlogger(testFile)
	for i := 0; i < b.N; i++ {
		g.log(ERROR, "Success", g, 1, true)
	}
	g.Cleanup()
	os.RemoveAll(testFile)
}
