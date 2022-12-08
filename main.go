package main

import (
	"os"
	"time"

	"github.com/jkamenik/advent-of-code-golang/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Setup a same message format
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})

	log.Info().Msg("Default Logger setup")

	cmd.Execute()
}
