package logger

type LogInterface interface {
    Emer(format string, v ...interface{})
    Alert(format string, v ...interface{})
    Crit(format string, v ...interface{})
    Error(format string, v ...interface{})
    Warn(format string, v ...interface{})
    Debug(format string, v ...interface{})
    Trace(format string, v ...interface{})
    Info(format string, v ...interface{})
    Panic(format string, v ...interface{})
    Fatal(format string, v ...interface{})
}

type LogNil struct {
    LogInterface
}

func NewLogNil() *LogNil {
    return new(LogNil)
}
func (l *LogNil) Emer(format string, v ...interface{}) {
}
func (l *LogNil) Panic(format string, v ...interface{}) {}
func (l *LogNil) Fatal(format string, v ...interface{}) {}
func (l *LogNil) Alert(format string, v ...interface{}) {
}
func (l *LogNil) Error(format string, v ...interface{}) {
}
func (l *LogNil) Warn(format string, v ...interface{}) {
}
func (l *LogNil) Debug(format string, v ...interface{}) {
}
func (l *LogNil) Trace(format string, v ...interface{}) {
}
func (l *LogNil) Info(format string, v ...interface{}) {
}

type GtLogger struct {
    log LogInterface
}

func NewGtLogger(log LogInterface) *GtLogger {
    l := new(GtLogger)
    l.log = log
    return l
}

// Painc logs a message at emergency level and panic.
func (defaultLogger *GtLogger) Painc(f interface{}, v ...interface{}) {
    defaultLogger.log.Panic(formatLog(f, v...))
}

// Fatal logs a message at emergency level and exit.
func (defaultLogger *GtLogger) Fatal(f interface{}, v ...interface{}) {
    defaultLogger.log.Fatal(formatLog(f, v...))
}

// Emer logs a message at emergency level.
func (defaultLogger *GtLogger) Emer(f interface{}, v ...interface{}) {
    defaultLogger.log.Emer(formatLog(f, v...))
}

// Alert logs a message at alert level.
func (defaultLogger *GtLogger) Alert(f interface{}, v ...interface{}) {
    defaultLogger.log.Alert(formatLog(f, v...))
}

// Crit logs a message at critical level.
func (defaultLogger *GtLogger) Crit(f interface{}, v ...interface{}) {
    defaultLogger.log.Crit(formatLog(f, v...))
}

// Error logs a message at error level.
func (defaultLogger *GtLogger) Error(f interface{}, v ...interface{}) {
    defaultLogger.log.Error(formatLog(f, v...))
}

// Warn logs a message at warning level.
func (defaultLogger *GtLogger) Warn(f interface{}, v ...interface{}) {
    defaultLogger.log.Warn(formatLog(f, v...))
}

// Info logs a message at info level.
func (defaultLogger *GtLogger) Info(f interface{}, v ...interface{}) {
    defaultLogger.log.Info(formatLog(f, v...))
}

// Notice logs a message at debug level.
func (defaultLogger *GtLogger) Debug(f interface{}, v ...interface{}) {
    defaultLogger.log.Debug(formatLog(f, v...))
}

// Trace logs a message at trace level.
func (defaultLogger *GtLogger) Trace(f interface{}, v ...interface{}) {
    defaultLogger.log.Trace(formatLog(f, v...))
}
