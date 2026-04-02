package main

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Info().
		Str("user", "aditya").
		Int("id", 101).
		Msg("User logged in")

	logger.Error().
		Msg("Something failed")
}
