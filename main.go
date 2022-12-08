package main

import (
	"github.com/jkamenik/advent-of-code-golang/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"time"
	"strings"
	"fmt"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Setup a same message format
	log.Logger = log.Output(zerolog.ConsoleWriter{
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("%s |", i)
		},
	})

	cmd.Execute()
}
