package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	logger.Info("User logged in",
		zap.String("user", "aditya"),
		zap.Int("id", 101),
	)

	logger.Error("Something failed",
		zap.String("error", "db connection"),
	)
}
