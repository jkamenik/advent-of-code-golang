package twentyTwentyTwo

import (
	"fmt"
	"os"

	"github.com/jkamenik/advent-of-code-golang/input"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var d1p1 = &cobra.Command{
	Use:   "d1p1 <input file>",
	Short: "day 1 puzzle 1",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Day 1 Puzzle 1")

		stream, err := input.StreamFile(args[0], 1)
		if err != nil {
			log.Err(err).Msgf("Unable to read file %v", args[0])
			os.Exit(1)
		}

		// Convert to Ints
		ints := IntChanAsArray(
			reduce(
				input.StringChanToIntChan(stream)))

		log.Info().Str("ints", fmt.Sprintf("%v", ints)).Msg("The array")
	},
}

func init() {
	twentyTwentyTwo.AddCommand(d1p1)
}

func IntChanAsArray(in <-chan int) []int {
	out := make([]int, 0)

	for i := range in {
		out = append(out, i)
	}

	return out
}
