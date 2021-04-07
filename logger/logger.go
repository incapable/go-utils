package logger

import (
	"git.poofycow.com/mark/go-utils/program"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func init() {
	initLogger()
}

var GlobalLogger *zap.Logger

// global logger configuration
func setupLogger(debug bool) {
	var (
		toFile = program.ReadProgramFlag("logToFile")

		outFilter = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.InfoLevel || (debug && lvl == zapcore.DebugLevel)
		})
		errFilter = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl > zapcore.InfoLevel
		})

		stdOut = zapcore.Lock(os.Stdout)
		stdErr = zapcore.Lock(os.Stderr)
	)

	var encoderConfig zapcore.EncoderConfig

	if !debug {
		encoderConfig = zap.NewProductionEncoderConfig()
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}

	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.LevelKey = "level"
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.MessageKey = "message"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder

	console := zapcore.NewConsoleEncoder(encoderConfig)

	if toFile {
		file, err := os.Create("log")
		if err != nil {
			panic("Could not create log file")
		}

		GlobalLogger = zap.New(zapcore.NewTee(
			zapcore.NewCore(console, zapcore.AddSync(stdOut), outFilter),
			zapcore.NewCore(console, zapcore.AddSync(stdErr), errFilter),
			zapcore.NewCore(console, zapcore.AddSync(file), zap.DebugLevel),
		))
	} else {
		GlobalLogger = zap.New(zapcore.NewTee(
			zapcore.NewCore(console, zapcore.AddSync(stdOut), outFilter),
			zapcore.NewCore(console, zapcore.AddSync(stdErr), errFilter),
		))
	}
}
