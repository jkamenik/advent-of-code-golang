package main

import (
	"github.com/rs/zerolog"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.Infolevel)

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
}
