package twentyTwentyTwo

import (
	"fmt"
	"os"
	"sort"

	"github.com/jkamenik/advent-of-code-golang/input"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var d1p2 = &cobra.Command{
	Use:   "d1p2 <input file>",
	Short: "day 1 puzzle 2",
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

		sort.Ints(ints)
		log.Debug().Msgf("Sorted: %v", ints)

		top := ints[len(ints)-3 : len(ints)]

		fmt.Printf("Max values are: %v\n", top)

		sum := 0
		for _, x := range top {
			sum = sum + x
		}

		fmt.Printf("Sum of the top 3: %v\n", sum)
	},
}

func init() {
	twentyTwentyTwo.AddCommand(d1p2)
}
