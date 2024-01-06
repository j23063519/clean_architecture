package log

import (
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

// NewLogger new logger
//
// filename path:storage/logs/logs.log
//
// maxSize maximum storage Unit: MB
//
// maxBackup maximum amount of file，0 no limited，but if maxDay then still delete the file
//
// maxDay The maximum number of days to save, 7 means the logs from one week ago will be deleted, 0 means not deleted
//
// compress default: false，cause if application goes wrong then we can read the logs
//
// logType single: unique file、daily: single file per day
//
// level
//
// [debug]: at developing ex: http、database request、send mail、sms
//
// [info]: at service ex: login、order
//
// [warn]: if you want to check something
//
// [error]: at error in production ex: database connection error、panic、http error
//
// The above log levels are from low to high. The higher the log level, the fewer messages are recorded.
func NewLogger(filename string, maxSize, maxBackup, maxDay int, compress bool, logType string, level string) {
	// logging method
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxDay, compress, logType)

	// set log level
	logLevel := new(zapcore.Level)
	if logLevel.UnmarshalText([]byte(level)) != nil {
		logLevel.UnmarshalText([]byte("debug"))
	}

	// new zapcore
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)

	// new Logger
	Logger = zap.New(core,
		zap.AddCaller(),                   // internal use runtime.Caller
		zap.AddCallerSkip(1),              // because runtime.Caller(1) then need skip i
		zap.AddStacktrace(zap.ErrorLevel), // if error then will show stacktrace
	)

	// transform custom logger to global logger
	// when use zap.L().Fatal() then use custom Logger
	zap.ReplaceGlobals(Logger)
}

// set logger format
func getEncoder() zapcore.Encoder {
	// format and rule
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // add "\n" in the end of line
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // use uppercase in log level，ex: ERROR INFO
		EncodeTime:     customTimeEncoder,              // date time: 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // excution time，Unit: second
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller relative path
	}

	// JSON encoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// customTimeEncoder
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// log record style: document and terminal
func getLogWriter(filename string, maxSize, maxBackup, maxDay int, compress bool, logType string) zapcore.WriteSyncer {
	if logType == "daily" {
		logname := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename, "logs.log", logname)
	}

	// rolling log
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxDay,
		Compress:   compress,
	}

	return zapcore.AddSync(lumberJackLogger)
}
