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

type logger struct {
	logLvl int
}

func NewLogger() ILogger {
	return &logger{}
}

func (l *logger) Log(lvl int, arg any) {
	if l.logLvl&lvl == 0 {
		return
	}

	switch lvl {
	case LogLvlDebug:
		println(Greyf("DEBUG: %s", arg))
	case LogLvlInfo:
		println(Whitef("INFO: %s", arg))
	case LogLvlWarn:
		println(Yellowf("WARN: %s", arg))
	case LogLvlError:
		println(Redf("ERROR: %s", arg))
	case LogLvlFatal:
		println(RedBoldf("FATAL: %s", arg))
	}
}

func (l *logger) LogLvl(lvl int) {
	l.logLvl = lvl
}
