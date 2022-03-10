package glog

import (
	"os"
	"testing"
)

const testFile = "test.log"

func TestUnstructureGlogger(t *testing.T) {
	type args struct {
		filePath string
		options  []Options
	}

	tests := []struct {
		name    string
		args    args
		wantNil bool
	}{
		{
			name: "success",
			args: args{
				filePath:testFile,
				options:  []Options{},
			},
			wantNil: false,
		},
		{
			name: "fail",
			args: args{
				filePath: "unknown/path/test.log",
				options:  []Options{},
			},
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUnstructureGlogger(tt.args.filePath, tt.args.options...); !tt.wantNil && (got == nil) {
				t.Errorf("UnstructureGlogger() = %v, want %v", got, "valid logger")
			}
		})
	}
	os.RemoveAll(testFile)
}

func TestJSONGlogger(t *testing.T) {
	type args struct {
		filePath string
		options  []Options
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
	}{
		{
			name: "success",
			args: args{
				filePath:testFile,
				options:  []Options{},
			},
			wantNil: false,
		},
		{
			name: "fail",
			args: args{
				filePath: "unknown/path/test.log",
				options:  []Options{},
			},
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJSONGlogger(tt.args.filePath, tt.args.options...); !tt.wantNil && (got == nil) {
				t.Errorf("JSONGlogger() = %v, want %v", got, "valid logger")
			}
		})
	}
}

func TestLogLevel_string(t *testing.T) {
	tests := []struct {
		name  string
		level LogLevel
		want  string
	}{
		{
			name:  "Info",
			level: INFO,
			want:  "INFO",
		},
		{
			name:  "Debug",
			level: DEBUG,
			want:  "DEBUG",
		},
		{
			name:  "Error",
			level: ERROR,
			want:  "ERROR",
		},
		{
			name:  "Warn",
			level: WARN,
			want:  "WARN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.level.string(); got != tt.want {
				t.Errorf("LogLevel.string() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultGlog_log(t *testing.T) {
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
			g:    defaultGlogger,
			args: args{
				format: "%s",
				level:  ERROR,
				msg:    []interface{}{"Success"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.msg...)
			Info(tt.args.msg...)
			Warn(tt.args.msg...)
			Debug(tt.args.msg...)
		})
	}
	os.RemoveAll(testFile)
}

func TestDefaultGlog_logf(t *testing.T) {
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
			g:    defaultGlogger,
			args: args{
				format: "%s",
				level:  ERROR,
				msg:    []interface{}{"Success"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorf(tt.args.format, tt.args.msg...)
			Infof(tt.args.format, tt.args.msg...)
			Warnf(tt.args.format, tt.args.msg...)
			Debugf(tt.args.format, tt.args.msg...)
		})
	}
	os.RemoveAll(testFile)
}

// func TestCleanup(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		{
// 			name: "cleanup",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			defaultGlogger.Info("cleanup")
// 			Cleanup()
// 		})
// 	}
// }
