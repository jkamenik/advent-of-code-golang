package twentyTwentyTwo

import (
	"fmt"
	"os"
	"strings"

	"github.com/jkamenik/advent-of-code-golang/input"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var d2p1 = &cobra.Command{
	Use:   "d2p1 <input file>",
	Short: "day 2 puzzle 1",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stream, err := input.StreamFile(args[0], 1)
		if err != nil {
			log.Err(err).Msgf("Unable to read file %v", args[0])
			os.Exit(1)
		}

		games := RPSGame(stream)

		p1Score := 0
		p2Score := 0
		for game := range games {
			x := game.P1Score()
			y := game.P2Score()

			log.Info().Str("game", game.String()).Int("P1", x).Int("P2", y).Msg("Scores")

			p1Score = p1Score + x
			p2Score = p2Score + y
		}

		fmt.Printf("Total Score are: P1=%v, P2=%v\n", p1Score, p2Score)
	},
}

func init() {
	twentyTwentyTwo.AddCommand(d2p1)
}

// RPSGame takes an input channel and makes an output of RPS games
func RPSGame(in <-chan string) <-chan RPS {
	out := make(chan RPS)

	go func() {
		defer close(out)

		for move := range in {
			plays := strings.Split(move, " ")

			if len(plays) < 2 {
				log.Warn().Msgf("Unexpected number of plays for %v", plays)
				continue
			}

			out <- RPS{plays[0], plays[1]}
		}
	}()

	return out
}

var RPSLookup = map[string][]string{
	"rock":     []string{"A", "X"},
	"paper":    []string{"B", "Y"},
	"scissors": []string{"C", "Z"},
}
var RPSBaseScore = map[string]int{
	"A": 1, "X": 1, // Rock
	"B": 2, "Y": 2, // Paper
	"C": 3, "Z": 3, // Scissors
}
var RPSTranspose = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
	"X": "A",
	"Y": "B",
	"Z": "C",
}

type RPS struct {
	p1 string
	p2 string
}

func (r RPS) String() string {
	return fmt.Sprintf("RPS{%v, %v}", r.p1, r.p2)
}

// P1Score calculates the score based on the following rules
// 1. The value of your choice (1 for rock, 2 for paper, etc...)
// 2. The outcome of the round (0 if you lost, 3 for a draw, 6 for a win)
func (r RPS) P1Score() int {
	base := RPSBaseScore[r.p1]
	p2 := RPSTranspose[r.p2]

	log.Debug().Int("base", base).Str("normalizedP2", p2).Str("game", r.String()).Msg("")

	// Draw
	if r.p1 == p2 {
		log.Info().Msg("Draw")
		return base + 3
	}

	// Rock beats Scissors
	if r.p1 == "A" && p2 == "C" {
		log.Info().Msg("Rock beats Scissors")
		return base + 6
	}

	// Paper beats Rock
	if r.p1 == "B" && p2 == "A" {
		log.Info().Msg("Paper beats Rock")
		return base + 6
	}

	// Scissors beats Paper
	if r.p1 == "C" && p2 == "B" {
		log.Info().Msg("Scissors beats Paper")
		return base + 6
	}

	log.Info().Msg("A loss")
	return base
}

func (r RPS) P2Score() int {
	base := RPSBaseScore[r.p1]
	p2 := RPSTranspose[r.p2]

	log.Debug().Int("base", base).Str("normalizedP2", p2).Str("game", r.String()).Msg("")

	// Draw
	if r.p1 == p2 {
		log.Info().Msg("Draw")
		return base + 3
	}

	// Rock beats Scissors
	if p2 == "A" && r.p1 == "C" {
		log.Info().Msg("Rock beats Scissors")
		return base + 6
	}

	// Paper beats Rock
	if p2 == "B" && r.p1 == "A" {
		log.Info().Msg("Paper beats Rock")
		return base + 6
	}

	// Scissors beats Paper
	if p2 == "C" && r.p1 == "B" {
		log.Info().Msg("Scissors beats Paper")
		return base + 6
	}

	log.Info().Msg("A loss")
	return base
}
