package twentyTwentyThree

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jkamenik/advent-of-code-golang/cmd/2023/day5"
	"github.com/jkamenik/advent-of-code-golang/input"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Puzzle is a function that takes the file input and returns the answer
// or an error
type Puzzle func(filename string, file <-chan string) (output string, err error)

// RegisterPuzzle registers a puzzle for execution
func RegisterPuzzle(day string, puzzle Puzzle) {
	log.Trace().Str("lookup", day).Msg("Registered puzzle")
	puzzleLookup[day] = puzzle
}

var puzzleLookup = map[string]Puzzle{}

var twentyTwentyThree = &cobra.Command{
	Use:   "2023 <day> <puzzle> [<file>]",
	Short: "2023 questions",
	Long: `Run the day and puzzle in question and provides the output.

The input file is assumed to be in ./advent-questions/2023/<day>.txt.  To change this provide the file path as the last argument

Output is provided to the screen.
`,
	Args: cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		lookup := fmt.Sprintf("%s-%s", args[0], args[1])
		file := ""

		if len(args) > 2 {
			file = args[2]
		} else {
			currentDir, err := os.Getwd()
			if err != nil {
				return err
			}

			file = filepath.Join(currentDir, "advent-questions", "2023", fmt.Sprintf("%s.txt", args[0]))
		}

		p, ok := puzzleLookup[lookup]
		if !ok {
			return fmt.Errorf("Puzzle %s could not be found", lookup)
		}

		stream, err := input.StreamFile(file, 10)
		if err != nil {
			return err
		}

		output, err := p(file, stream)
		if err != nil {
			return err
		}

		fmt.Println(output)

		return nil
	},
}

// Load adds our root command to the command we are given
func Load(root *cobra.Command) {
	root.AddCommand(twentyTwentyThree)

	// Now connect all the puzzles
	RegisterPuzzle("1-1", d1p1)
	RegisterPuzzle("1-2", d1p2)
	RegisterPuzzle("2-1", d2p1)
	RegisterPuzzle("2-2", d2p2)
	RegisterPuzzle("3-1", d3p1)
	RegisterPuzzle("3-2", d3p2)
	RegisterPuzzle("4-1", d4p1)
	RegisterPuzzle("4-2", d4p2)
	RegisterPuzzle("5-1", day5.Part1)
	RegisterPuzzle("5-2", day5.Part2)
}
