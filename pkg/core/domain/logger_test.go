package domain

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name    string
		wantOut string
	}{
		{
			name:    "Test Empty NewLogger",
			wantOut: "",
		},
		{
			name:    "Test Non-Empty NewLogger",
			wantOut: "Test message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := new(bytes.Buffer)
			gotLogger := NewLogger(out)

			// Assuming NewLogger should return a non-nil logger
			if gotLogger == nil {
				t.Error("NewLogger() = nil, want non-nil")
			}

			// Write to the logger if wantOut is not empty
			if tt.wantOut != "" {
				gotLogger.Info(tt.wantOut) // Assuming you have an Info method to write logs
				if gotOut := out.String(); !strings.Contains(gotOut, tt.wantOut) {
					t.Errorf("Logger output = %v, want to contain %s", gotOut, tt.wantOut)
				}
			}
		})
	}
}

func TestLogger_Info(t *testing.T) {
	type fields struct {
		infoLog  *log.Logger
		warnLog  *log.Logger
		errorLog *log.Logger
		fatalLog *log.Logger
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Test Full with Message",
			fields: fields{
				infoLog:  log.New(new(bytes.Buffer), "INFO: ", log.Ldate|log.Ltime),
				warnLog:  log.New(new(bytes.Buffer), "WARN: ", log.Ldate|log.Ltime),
				errorLog: log.New(new(bytes.Buffer), "ERROR: ", log.Ldate|log.Ltime),
				fatalLog: log.New(new(bytes.Buffer), "FATAL: ", log.Ldate|log.Ltime),
			},
			args: args{msg: "Test Full with Empty Message"},
		},
		{name: "Test Empty Info",
			fields: fields{
				infoLog:  log.New(new(bytes.Buffer), "INFO: ", log.Ldate|log.Ltime),
				warnLog:  log.New(new(bytes.Buffer), "WARN: ", log.Ldate|log.Ltime),
				errorLog: log.New(new(bytes.Buffer), "ERROR: ", log.Ldate|log.Ltime),
				fatalLog: log.New(new(bytes.Buffer), "FATAL: ", log.Ldate|log.Ltime),
			},
			args: args{msg: ""},
		},
		{name: "Test Nil with Empty Message",
			fields: fields{
				infoLog:  nil,
				warnLog:  nil,
				errorLog: nil,
				fatalLog: nil,
			},
			args: args{msg: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				infoLog:  tt.fields.infoLog,
				warnLog:  tt.fields.warnLog,
				errorLog: tt.fields.errorLog,
				fatalLog: tt.fields.fatalLog,
			}
			l.Info(tt.args.msg)
			l.Warn(tt.args.msg)
			//go l.Info(tt.args.msg)
			//go l.Warn(tt.args.msg)
			//go l.Error(tt.args.msg)
			//go l.Fatal(tt.args.msg)
		})
	}
}
