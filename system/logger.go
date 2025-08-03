package system

import "github.com/antigloss/go/logger"

var Logger *logger.Logger

func InitLogger(prefix string) {
	var err error
	Logger, err = logger.New(&logger.Config{
		LogFileMaxSize:    25,
		LogFileMaxNum:     100,
		LogFileNumToDel:   50,
		LogFilenamePrefix: prefix,
		LogDir:            Env("SYS_LOG_DIR", "./logs"),
		LogLevel:          logger.LogLevel(EnvInt("SYS_LOG_LEVEL", int(logger.LogLevelWarn))),
		LogDest:           logger.LogDest(EnvInt("SYS_LOG_DEST", int(logger.LogDestBoth))),
		Flag:              logger.ControlFlagLogDate | logger.ControlFlagLogLineNum,
	})
	if err != nil {
		panic(err)
	}
}
