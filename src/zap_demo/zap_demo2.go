package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)


func main() {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	outFile,_:=os.Create("./logs/out.log")
	topicDebugging := zapcore.AddSync(outFile)
	errorFile,_:=os.Create("./logs/error.log")
	topicErrors := zapcore.AddSync(errorFile)

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),

		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)
	logger := zap.New(core)
	defer logger.Sync()

	logger.Info("constructed a logger Info")
	logger.Error("constructed a logger Error")

	logger.Sugar().Info("Sugar constructed a logger Info")
	logger.Sugar().Error("Sugar constructed a logger Error")
}


