package logger

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rs/zerolog"
)

const TIME_FORMAT = "2006-01-02T15:04:05.000000"

var (
	AppLog        *rotatelogs.RotateLogs
	MiddlewareLog *rotatelogs.RotateLogs
	config        zerolog.Logger
)

func init() {
	godotenv.Load(".env")
	var (
		LogDir    = os.Getenv("LOG_DIR")
		LogMaxAge = os.Getenv("LOG_MAX_AGE")
		Debug     = os.Getenv("LOG_DEBUG")
	)
	currentDir, _ := os.Getwd()
	LogDir = currentDir + LogDir
	MaxAge, _ := strconv.Atoi(LogMaxAge)

	if MaxAge < 1 {
		MaxAge = 15
	}

	AppLog, _ = rotatelogs.New(
		LogDir+"/serv_log.%Y%m%d%H%M",
		rotatelogs.WithLinkName(LogDir+"/serv_log"),
		rotatelogs.WithMaxAge(time.Duration(MaxAge)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)

	MiddlewareLog, _ = rotatelogs.New(
		LogDir+"/http_log.%Y%m%d%H%M",
		rotatelogs.WithLinkName(LogDir+"/http_log"),
		rotatelogs.WithMaxAge(time.Duration(MaxAge)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)

	isDebug, _ := strconv.ParseBool(Debug)

	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)

	zerolog.TimeFieldFormat = TIME_FORMAT
	// change applog into stdout
	config = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func GetConfig() zerolog.Logger {
	return config
}
