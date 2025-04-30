package system

import (
	"fmt"
	"os"

	"github.com/ADM87/orion/system/logging"
)

var (
	logger = logging.NewLogger()
)

// ======================================================================
// Logging functions
// ======================================================================

func Debug(arg any) {
	logger.Log(logging.LogLvlDebug, arg)
}

func Debugf(format string, args any) {
	logger.Log(logging.LogLvlDebug, fmt.Sprintf(format, args))
}

func Info(arg any) {
	logger.Log(logging.LogLvlInfo, arg)
}

func Infof(format string, args any) {
	logger.Log(logging.LogLvlInfo, fmt.Sprintf(format, args))
}

func Warn(arg any) {
	logger.Log(logging.LogLvlWarn, arg)
}

func Warnf(format string, args any) {
	logger.Log(logging.LogLvlWarn, fmt.Sprintf(format, args))
}

func Error(arg any) {
	logger.Log(logging.LogLvlError, arg)
}

func Errorf(format string, args any) {
	logger.Log(logging.LogLvlError, fmt.Sprintf(format, args))
}

func Fatal(arg any) {
	logger.Log(logging.LogLvlFatal, arg)
	os.Exit(1)
}

func Fatalf(format string, args any) {
	logger.Log(logging.LogLvlFatal, fmt.Sprintf(format, args))
	os.Exit(1)
}
