package istio_logrus

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

/*
type logger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
}*/
/*
func New() *LogrusLogger {
	return &LogrusLogger{
		SkipErrRecordNotFound: true,
	}
}

func (l *logger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *logger) Info(ctx context.Context, s string, args ...interface{}) {
	log.WithContext(ctx).Infof(s, args)
}

func (l *logger) Warn(ctx context.Context, s string, args ...interface{}) {
	log.WithContext(ctx).Warnf(s, args)
}

func (l *logger) Error(ctx context.Context, s string, args ...interface{}) {
	log.WithContext(ctx).Errorf(s, args)
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := log.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[log.ErrorKey] = err
		log.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		log.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	log.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
}
*/
// ConsoleLogger is the struct used for mesh command
type LogrusLogger struct {
	/*stdOut io.Writer
	stdErr io.Writer
	scope  *log.Scope*/
}

// NewDefaultLogger creates a new logger that outputs to stdout/stderr at default scope.
/*func NewDefaultLogger() *ConsoleLogger {
	return NewConsoleLogger(os.Stdout, os.Stderr, nil)
}*/
func NewDefaultIstioLogrusLogger() *LogrusLogger {
	//return NewLogrusLogger(os.Stdout, os.Stderr, nil)
	return NewLogrusLogger()
}

// NewConsoleLogger creates a new logger and returns a pointer to it.
// stdOut and stdErr can be used to capture output for testing. If scope is nil, the default scope is used.
func NewLogrusLogger() *LogrusLogger {
	/*s := scope
	if s == nil {
		s = log.RegisterScope(log.DefaultScopeName, log.DefaultScopeName, 0)
	}
	return &LogrusLogger{
		stdOut: stdOut,
		stdErr: stdErr,
		scope:  s,
	}*/
	return &LogrusLogger{}
}

func (l *LogrusLogger) LogAndPrint(v ...interface{}) {
	if len(v) == 0 {
		return
	}
	/*s := fmt.Sprint(v...)
	l.Print(s + "\n")
	l.scope.Infof(s)*/
	s := fmt.Sprint(v...)
	log.Print(s + "\n")
	log.Infof(s)

}

func (l *LogrusLogger) LogAndError(v ...interface{}) {
	if len(v) == 0 {
		return
	}
	s := fmt.Sprint(v...)
	log.Errorf(s)
}

func (l *LogrusLogger) LogAndFatal(a ...interface{}) {
	l.LogAndError(a...)
	os.Exit(-1)
}

func (l *LogrusLogger) LogAndPrintf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Print(s + "\n")
	log.Infof(s)
}

func (l *LogrusLogger) LogAndErrorf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	log.Print(s + "\n")
	log.Errorf(s)
}

func (l *LogrusLogger) LogAndFatalf(format string, a ...interface{}) {
	l.LogAndErrorf(format, a...)
	os.Exit(-1)
}

func (l *LogrusLogger) Print(s string) {
	log.Print(s + "\n")
}

func (l *LogrusLogger) PrintErr(s string) {
	log.Print(s + "\n")
	log.Errorf(s)
}
