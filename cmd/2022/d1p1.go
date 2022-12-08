package twentyTwentyTwo

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var d1p1 = &cobra.Command{
	Use:   "d1p1 <input file>",
	Short: "day 1 puzzle 1",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Day 1 Puzzle 1")
	},
}

func init() {
	twentyTwentyTwo.AddCommand(d1p1)
}
