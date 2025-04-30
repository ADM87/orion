package logging

const (
	LogLvlDebug = iota
	LogLvlInfo
	LogLvlWarn
	LogLvlError
	LogLvlFatal
)

type ILogger interface {
	Log(lvl int, arg any)
}

type Logger struct {
}

func NewLogger() ILogger {
	return &Logger{}
}

func (l *Logger) Log(lvl int, arg any) {

}
