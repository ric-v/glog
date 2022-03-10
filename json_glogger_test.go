package glog

import (
	"os"
	"sync"
	"testing"
)

func TestJSONGlog_log(t *testing.T) {
	type args struct {
		format string
		level  LogLevel
		msg    []interface{}
	}
	tests := []struct {
		name string
		g    *JSONGlog
		args args
	}{
		{
			name: "success",
			g: &JSONGlog{
				glog: glog{
					out:   os.Stdout,
					queue: make(chan interface{}, 100),
					wg:    &sync.WaitGroup{},
				},
			},
			args: args{
				format: "%s",
				level:  ERROR,
				msg:    []interface{}{"Success", "message"},
			},
		},
		{
			name: "success with uneven json fields",
			g: &JSONGlog{
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

func TestJSONGlog_logf(t *testing.T) {
	type args struct {
		format string
		level  LogLevel
		msg    []interface{}
	}
	tests := []struct {
		name string
		g    *JSONGlog
		args args
	}{
		{
			name: "success",
			g: &JSONGlog{
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
func BenchmarkJSONGlog_logf(b *testing.B) {
	g := NewJSONGlogger(testFile)
	for i := 0; i < b.N; i++ {
		g.logf("", INFO, "key", "value", "key-2", 10, "key-3", true, "key-4", g)
	}
	g.Cleanup()
	os.RemoveAll(testFile)
}

// benchmarking
func BenchmarkJSONGlog_log(b *testing.B) {
	g := NewJSONGlogger(testFile)
	for i := 0; i < b.N; i++ {
		g.log(INFO, "key", "value", "key-2", 10, "key-3", true, "key-4", g)
	}
	g.Cleanup()
	os.RemoveAll(testFile)
}
