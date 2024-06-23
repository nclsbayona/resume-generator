package domain

import (
	"io"
	"log"
)

type Logger struct {
	infoLog  *log.Logger
	warnLog  *log.Logger
	errorLog *log.Logger
	fatalLog *log.Logger
}

func NewLogger(out io.Writer) (logger *Logger) {
	logger = &Logger{
		infoLog:  log.New(out, "INFO: ", log.Ldate|log.Ltime),
		warnLog:  log.New(out, "WARN: ", log.Ldate|log.Ltime),
		errorLog: log.New(out, "ERROR: ", log.Ldate|log.Ltime),
		fatalLog: log.New(out, "FATAL: ", log.Ldate|log.Ltime),
	}
	return
}

func (l *Logger) Info(msg string) {
	if l.infoLog != nil {
		l.infoLog.Println(msg)
	}
}

func (l *Logger) Warn(msg string) {
	if l.infoLog != nil {
		l.warnLog.Println(msg)
	}
}

func (l *Logger) Error(msg string) {
	if l.infoLog != nil {
		l.errorLog.Panic(msg)
	}
}

func (l *Logger) Fatal(msg string) {
	if l.infoLog != nil {
		l.fatalLog.Fatal(msg)
	}
}
