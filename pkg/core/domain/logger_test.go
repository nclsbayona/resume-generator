package domain

import (
	"bytes"
	"log"
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {
	name := "Test Empty NewLogger"
	out := new(bytes.Buffer)
	wantLogger := &Logger{
		infoLog:  log.New(out, "INFO: ", log.Ldate|log.Ltime),
		warnLog:  log.New(out, "WARN: ", log.Ldate|log.Ltime),
		errorLog: log.New(out, "ERROR: ", log.Ldate|log.Ltime),
		fatalLog: log.New(out, "FATAL: ", log.Ldate|log.Ltime),
	}

	type test struct {
		name       *string
		wantLogger *Logger
		wantOut    *string
	}

	wantOut := out.String()

	testCase := test{
		name: &name, wantLogger: wantLogger, wantOut: &wantOut,
	}

	t.Run(*testCase.name, func(t *testing.T) {
		gotLogger := NewLogger(out)
		if !reflect.DeepEqual(gotLogger, testCase.wantLogger) {
			t.Errorf("NewLogger() = %v, want %v", gotLogger, testCase.wantLogger)
		}
		if gotOut := out.String(); gotOut != *testCase.wantOut {
			t.Errorf("NewLogger() = %v, want %s", gotOut, *testCase.wantOut)
		}
	})

	name = "Test Non-Empty NewLogger"

	out = new(bytes.Buffer)

	wantLogger = &Logger{
		infoLog:  log.New(out, "INFO: ", log.Ldate|log.Ltime),
		warnLog:  log.New(out, "WARN: ", log.Ldate|log.Ltime),
		errorLog: log.New(out, "ERROR: ", log.Ldate|log.Ltime),
		fatalLog: log.New(out, "FATAL: ", log.Ldate|log.Ltime),
	}

	wantOut = "Test message"

	testCase = test{
		name: &name, wantLogger: wantLogger, wantOut: &wantOut,
	}

	t.Run(*testCase.name, func(t *testing.T) {
		gotLogger := NewLogger(out)
		if !reflect.DeepEqual(gotLogger, testCase.wantLogger) {
			t.Errorf("NewLogger() = %v, want %v", gotLogger, testCase.wantLogger)
		}
		out.WriteString(wantOut)
		if gotOut := out.String(); gotOut != *testCase.wantOut {
			t.Errorf("NewLogger() = %v, want %s", gotOut, *testCase.wantOut)
		}
	})
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
