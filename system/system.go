package system

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ADM87/orion/system/logging"
)

var (
	logger = logging.NewLogger()
)

func init() {
	logger.LogLvl(logging.LogLvlError | logging.LogLvlFatal)
}

// ======================================================================
// System functions
// ======================================================================

// MakeAbsolute takes a path and returns the absolute path. If the path is already absolute, it returns the same path.
func MakeAbsolute(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

// ======================================================================
// Logging functions
// ======================================================================

// LogLvl sets the logging level for the logger
func LogLvl(lvl int) {
	logger.LogLvl(lvl)
}

// Debug logs debug level messages
func Debug(arg any) {
	logger.Log(logging.LogLvlDebug, arg)
}

// Debugf logs debug level messages with formatting
func Debugf(format string, args any) {
	logger.Log(logging.LogLvlDebug, fmt.Sprintf(format, args))
}

// Info logs info level messages
func Info(arg any) {
	logger.Log(logging.LogLvlInfo, arg)
}

// Infof logs info level messages with formatting
func Infof(format string, args any) {
	logger.Log(logging.LogLvlInfo, fmt.Sprintf(format, args))
}

// Warn logs warning level messages
func Warn(arg any) {
	logger.Log(logging.LogLvlWarn, arg)
}

// Warnf logs warning level messages with formatting
func Warnf(format string, args any) {
	logger.Log(logging.LogLvlWarn, fmt.Sprintf(format, args))
}

// Error logs error level messages
func Error(arg any) {
	logger.Log(logging.LogLvlError, arg)
}

// Errorf logs error level messages with formatting
func Errorf(format string, args any) {
	logger.Log(logging.LogLvlError, fmt.Sprintf(format, args))
}

// Fatal logs fatal level messages, which will also exit the program
func Fatal(arg any) {
	logger.Log(logging.LogLvlFatal, arg)
	os.Exit(1)
}

// Fatalf logs fatal level messages with formatting, which will also exit the program
func Fatalf(format string, args any) {
	logger.Log(logging.LogLvlFatal, fmt.Sprintf(format, args))
	os.Exit(1)
}
