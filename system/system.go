package system

import (
	"fmt"
	"os"
	"path/filepath"

	systemargs "github.com/ADM87/orion/system/args"
	systemlog "github.com/ADM87/orion/system/log"
)

var (
	Verbose = systemargs.NewBoolArg("verbose", "v", "Enable verbose logging", false)
)

var (
	logger = systemlog.NewLogger()
)

func init() {
	logger.LogLvl(systemlog.LogLvlError | systemlog.LogLvlFatal)
}

// ======================================================================
// System functions
// ======================================================================

func Exit(code int) {
	Debugf("Exiting with code %d", code)
	os.Exit(code)
}

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

// SetVerbose sets the logging level based on the Verbose argument
func SetVerbose() {
	if Verbose.GetValue() {
		logger.LogLvl(systemlog.LogLvlAll)
	} else {
		logger.LogLvl(systemlog.LogLvlError | systemlog.LogLvlFatal)
	}
}

// LogLvl sets the logging level for the logger
func LogLvl(lvl int) {
	logger.LogLvl(lvl)
}

// Debug logs debug level messages
func Debug(arg any) {
	logger.Log(systemlog.LogLvlDebug, arg)
}

// Debugf logs debug level messages with formatting
func Debugf(format string, args any) {
	logger.Log(systemlog.LogLvlDebug, fmt.Sprintf(format, args))
}

// Info logs info level messages
func Info(arg any) {
	logger.Log(systemlog.LogLvlInfo, arg)
}

// Infof logs info level messages with formatting
func Infof(format string, args any) {
	logger.Log(systemlog.LogLvlInfo, fmt.Sprintf(format, args))
}

// Warn logs warning level messages
func Warn(arg any) {
	logger.Log(systemlog.LogLvlWarn, arg)
}

// Warnf logs warning level messages with formatting
func Warnf(format string, args any) {
	logger.Log(systemlog.LogLvlWarn, fmt.Sprintf(format, args))
}

// Error logs error level messages
func Error(arg any) {
	logger.Log(systemlog.LogLvlError, arg)
}

// Errorf logs error level messages with formatting
func Errorf(format string, args any) {
	logger.Log(systemlog.LogLvlError, fmt.Sprintf(format, args))
}

// Fatal logs fatal level messages, which will also exit the program
func Fatal(arg any) {
	logger.Log(systemlog.LogLvlFatal, arg)
	os.Exit(1)
}

// Fatalf logs fatal level messages with formatting, which will also exit the program
func Fatalf(format string, args any) {
	logger.Log(systemlog.LogLvlFatal, fmt.Sprintf(format, args))
	os.Exit(1)
}
