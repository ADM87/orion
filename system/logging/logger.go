package logging

const (
	LogLvlDebug = 1 << iota
	LogLvlInfo
	LogLvlWarn
	LogLvlError
	LogLvlFatal

	LogLvlNone = 0
	LogLvlAll  = LogLvlDebug | LogLvlInfo | LogLvlWarn | LogLvlError | LogLvlFatal
)

type ILogger interface {
	Log(lvl int, arg any)
	LogLvl(lvl int)
}

type Logger struct {
	logLvl int
}

func NewLogger() ILogger {
	return &Logger{}
}

func (l *Logger) Log(lvl int, arg any) {
	if l.logLvl&lvl != lvl {
		return
	}
}

func (l *Logger) LogLvl(lvl int) {
	l.logLvl = lvl
}
