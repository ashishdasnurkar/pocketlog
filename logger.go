package pocketlog

import "fmt"

// Logger is used to log certain information
type Logger struct {
	threshold Level
}

// New returns you a logger, ready to log at the required threshold.
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

// Debug formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

func (l *Logger) Infof(string format, args ...any) {

}

func (l *Logger) Errorf(string format, args ...any) {

}
