package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "-DEBUG: ", logger.Flags()),
		info:    log.New(writer, "-INFO: ", logger.Flags()),
		warning: log.New(writer, "-WARNING: ", logger.Flags()),
		err:     log.New(writer, "-ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Create Non-formatted Debug Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

// Create Non-formatted Info Logs
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

// Create Non-formatted Warning Logs
func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

// Create Non-formatted Error Logs
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Create Formatted Debug Logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

// Create Formatted Info Logs
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

// Create Formatted Warning Logs
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

// Create Formatted Error Logs
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
