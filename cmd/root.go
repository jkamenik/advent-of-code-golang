package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var verbosity = 0

var rootCmd = &cobra.Command{
	Use:     "advent-of-code-golang",
	Short:   "Advent of code solutions",
	Version: "2022",
	PersistentPreRunE: func(cmd *cobra.Command, arg []string) error {
		err := setup()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().CountVarP(&verbosity, "verbose", "v", "More v's = more verbosity")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func setup() {
	log.Trace().Msg("Setting log verbosity")
	if verbosity >= 1 {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if verbosity >= 2 {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	log.Debug().Msgf("Verbosity: %d", verbosity)
	return nil
}
