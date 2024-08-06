package pocketlog

// Level represents the available level of logging
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information deemed valuable.
	LevelInfo
	// LevelError represents the highest logging level, only to be used to trace. errors
	LevelError
)

func (l *Logger) Debugf(string format, args ...any) {

}
func (l *Logger) Infof(string format, args ...any) {

}
