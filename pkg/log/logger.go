package log

import (
	"os"
	"sync"

	_logger "github.com/sirupsen/logrus"
)

var l Logger

type Logger struct {
	once     sync.Once
	_default *_logger.Logger
}

func (l *Logger) Init() {
	l.once.Do(func() {
		l._default = &_logger.Logger{
			Out:          os.Stderr,
			Formatter:    &_logger.TextFormatter{ForceColors: true},
			Hooks:        make(_logger.LevelHooks),
			Level:        _logger.TraceLevel,
			ExitFunc:     os.Exit,
			ReportCaller: false,
		}
	})
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	l.Init()
	l._default.Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	l.Init()
	l._default.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	l.Init()
	l._default.Print(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	l.Init()
	l._default.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	l.Init()
	l._default.Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	l.Init()
	l._default.Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	l.Init()
	l._default.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	l.Init()
	l._default.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	l.Init()
	l._default.Fatal(args...)
}
