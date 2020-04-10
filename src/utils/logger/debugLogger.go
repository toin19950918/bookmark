package logger

import (
	"fmt"
	"github.com/robin019/bookmark/src/utils/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"runtime"
	"time"
)

var (
	debugLog *zap.SugaredLogger
)

func Debug() *zap.SugaredLogger {
	if debugLog == nil {
		newDebugLogger()
	}
	return debugLog
}

func newDebugLogger() {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	logPath := fmt.Sprintf("%s/../../../%s%s", dir, config.Get("log.debugLoggerPath"), config.Get("log.debugLoggerFile"))

	logBuilder := zap.NewDevelopmentConfig()
	logBuilder.EncoderConfig.EncodeTime = SyslogTimeEncoder
	//logBuilder.EncoderConfig.EncodeLevel = CustomLevelEncoder
	logBuilder.OutputPaths = []string{
		logPath,
	}

	logger, err := logBuilder.Build()
	if err != nil {
		panic("can't initialize zap logger:" + err.Error())
	}
	debugLog = logger.Sugar()
}

func CloseDebugLogger() {
	if debugLog != nil {
		debugLog.Sync()
	}
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}
