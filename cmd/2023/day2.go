package twentyTwentyThree

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

func d2p1(filename string, file <-chan string)(string, error) {
	sum := int64(0)
	bag := d2Game{
		ID: 0,
		Red: 12,
		Green: 13,
		Blue: 14,
	}

	for line := range file {
		log.Trace().Msgf("line: %s", line)

		game, err := newD2Game(line)
		if err != nil {
			return "", err
		}
		log.Info().Msgf("game: %v -> %v", line, game)

		if bag.isPossible(game) {
			log.Debug().Msgf("Game is possible: %v", game)
			sum = sum + game.ID
		} else {
			log.Debug().Msgf("Game is NOT possible: %v", game)
		}
	}

	return fmt.Sprintf("%d", sum), nil
}

func d2p2(filename string, file <-chan string)(string, error) {
	sum := int64(0)

	for line := range file {
		log.Trace().Msgf("line: %s", line)

		game, err := newD2Game(line)
		if err != nil {
			return "", err
		}
		log.Info().Msgf("game: %v -> %v", line, game)

		sum = sum + game.Power()
	}

	return fmt.Sprintf("%d", sum), nil
}

type d2Game struct {
	ID int64
	Red int64
	Green int64
	Blue int64
}

func newD2Game(line string) (d2Game, error) {
	g := d2Game{}

	// Split on colon to get the digits
	items := strings.Split(line, ":")
	log.Trace().Msgf("%v", items)
	idStr := strings.Split(items[0], " ")[1]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return g, err
	}
	g.ID = id

	plays := strings.Split(items[1], ";")
	for _, play := range plays {
		log.Trace().Msgf("Play: %v", play)

		blocks := strings.Split(play, ",")
		for _, block := range blocks {
			log.Trace().Msgf("Block %s", block)

			split := strings.Split(block, " ")
			log.Trace().Msgf("Count: %v", split)
			number, err := strconv.ParseInt(split[1], 10, 64)
			if err != nil {
				return g, err
			}

			switch split[2] {
			case "red":
				if number > g.Red {
					g.Red = number
				}
			case "green":
				if number > g.Green {
					g.Green = number
				}
			case "blue":
				if number > g.Blue {
					g.Blue = number
				}
			}
		}
	}

	return g, nil
}

func (g d2Game) isPossible(o d2Game) bool {
	if o.Green <= g.Green &&
	   o.Red <= g.Red &&
	   o.Blue <= g.Blue {
		return true
	}

	return false
}

func (g d2Game) Power() int64 {
	return g.Red * g.Green * g.Blue
}


func init() {
	puzzleLookup["2-1"] = d2p1
	puzzleLookup["2-2"] = d2p2
}
