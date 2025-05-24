package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"sync"
)

type Config struct {

	// ConsoleLoggingEnabled console logging
	ConsoleLoggingEnabled bool

	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool

	// FileLoggingEnabled makes the framework log to a file.go
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool

	// Directory to log to when filelogging is enabled
	Directory string

	// Filename is the name of the logfile which will be placed inside the directory
	Filename string

	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int

	// MaxBackups the max number of rolled files to keep
	MaxBackups int

	// MaxAge the max age in days to keep a logfile
	MaxAge int

	// CallerSkip the number of directory hierarchy to be skipped
	CallerSkip int
}

type Logger struct {
	serviceLogger *zerolog.Logger
}

var (
	instance *Logger
	once     sync.Once
)

// InitZerolog, call it during cmd main.go
func InitZerolog(config Config) {
	once.Do(func() {
		instance = initializer(config)
	})
}

func initializer(config ...Config) *Logger {
	if len(config) == 0 {
		config = append(config, Config{
			ConsoleLoggingEnabled: true,
		})
	}
	var writers []io.Writer

	if config[0].ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if config[0].FileLoggingEnabled {
		writers = append(writers, newRollingFile(config[0]))
	}
	if config[0].CallerSkip == 0 {
		config[0].CallerSkip = 3
	}
	zerolog.CallerSkipFrameCount = config[0].CallerSkip
	mw := io.MultiWriter(writers...)

	logger := zerolog.New(mw).With().Timestamp().Caller().Logger()

	return &Logger{
		serviceLogger: &logger,
	}
}

func newRollingFile(config Config) io.Writer {
	l := &lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxBackups: config.MaxBackups, // files
		MaxSize:    config.MaxSize,    // megabytes
		MaxAge:     config.MaxAge,     // days
	}

	return l
}

// Call the singleton function
func Call() *Logger {
	if instance == nil {
		fmt.Println("logger not initialized, Call InitZerolog() first.")
		return nil
	}
	return instance
}

// Errorf print formatted error message
func (l *Logger) Errorf(format string, a ...interface{}) {
	errs := fmt.Errorf(format, a...)
	l.serviceLogger.Error().Msgf(errs.Error())
}

// Warnf print formatter warn message
func (l *Logger) Warnf(format string, a ...interface{}) {
	l.serviceLogger.Warn().Msgf(format, a...)
}

// Infof print formated info message
func (l *Logger) Infof(format string, a ...interface{}) {
	l.serviceLogger.Info().Msgf(format, a...)
}

// Fatalf
// Stop the app after invocation, and
// print Fatal message with format
func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.serviceLogger.Fatal().Msgf(format, a...)
}

// Panicf
// Stop the app after invocation,
// and print the Panic message
func (l *Logger) Panicf(format string, a ...interface{}) {
	l.serviceLogger.Panic().Msgf(format, a...)
}
