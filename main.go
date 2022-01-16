package main

import (
	"log"
	"os"
	"strconv"

	"github.com/dmatusiewicz/auth/cmd"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Setting up logging.
	logger := deployLogger()
	logger.Debug("Starting command engine.")
	err := cmd.Execute(logger)
	if err != nil {
		logger.Fatal(err.Error())
	}
}

//deployLogger simple function that create loger and sets the log level based on the zapcore.Level
func deployLogger() *zap.Logger {
	var defaultLogLevelS = "0"

	logLevelS, exists := os.LookupEnv("LOG_LEVEL")
	if !exists {
		logLevelS = defaultLogLevelS
	}

	logLevel, err := strconv.Atoi(logLevelS)
	if err != nil {
		log.Fatal(err)
	}

	atom := zap.NewAtomicLevel()
	atom.SetLevel(zapcore.Level(logLevel))
	enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	logger := zap.New(zapcore.NewCore(enc, zapcore.Lock(os.Stdout), atom))
	return logger
}
