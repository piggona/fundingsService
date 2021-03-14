package log

import "go.uber.org/zap"

// logger, _ := zap.NewProduction()
// defer logger.Sync() // flushes buffer, if any
// sugar := logger.Sugar()
// sugar.Infow("failed to fetch URL",
//   // Structured context as loosely typed key-value pairs.
//   "url", url,
//   "attempt", 3,
//   "backoff", time.Second,
// )
// sugar.Infof("Failed to fetch URL: %s", url)

func Info(template string, input ...interface{}) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Sugar().Infof(template, input)
}

func Warn(template string, input ...interface{}) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Sugar().Errorf(template, input)
}

func Error(template string, input ...interface{}) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Sugar().Errorf(template, input)
}
