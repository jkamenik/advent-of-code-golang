package twentyTwentyThree

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

func d4p1(filename string, file <-chan string) (string, error) {
	sum := 0

	for line := range file {
		card := NewD4Card(line)

		sum = sum + card.Score()
	}

	return fmt.Sprintf("%d", sum), nil
}

func d4p2(filename string, file <-chan string) (string, error) {
	return "", nil
}

type d4Card struct {
	ID      string
	Winners []string
	Numbers []string
}
func (c d4Card) String() string {
	return fmt.Sprintf("{id: '%v', w: %v, n: %v}", c.ID, c.Winners, c.Numbers)
}

func NewD4Card(line string) (card d4Card) {
	card.Winners = []string{}
	card.Numbers = []string{}

	matches := strings.Split(line, ":")
	log.Trace().Msgf("matches %v", matches)
	if len(matches) < 2 {
		return
	}

	id := DigitsFromString(matches[0])

	// id := strings.Split(matches[0], " ")
	log.Trace().Msgf("id part %v", id)
	card.ID = id[0]

	winners := strings.Split(matches[1], " | ")
	log.Trace().Msgf("winners and losers %+v", winners)

	card.Winners = DigitsFromString(winners[0])

	if len(winners) >= 2 {
		card.Numbers = DigitsFromString(winners[1])
	}

	return
}

func DigitsFromString(s string) []string {
	rtn := []string{}

	acc := ""

	for _, r := range s {
		log.Trace().Str("acc",acc).Str("digits",fmt.Sprintf("%v", rtn)).Msgf("rune: '%v'", string(r))

		if r >= '0' && r <= '9' {
			acc = acc + string(r)
		} else {
			log.Trace().Str("acc", acc).Msg("Not a number, flush accumulator")
			if acc != "" {
				rtn = append(rtn, acc)
				acc = ""
			}
		}
	}

	if acc != "" {
		rtn = append(rtn, acc)
	}

	log.Trace().Msgf("digits %v", rtn)

	return rtn
}

func (card d4Card) Score() int {
	score := 0

	for _, winner := range card.Winners {
		for _, have := range card.Numbers {
			if winner == have {
				log.Trace().Msgf("Found winner %s", winner)
				if score == 0 {
					score = 1
				} else {
					score = score * 2
				}
			}
		}
	}

	return score
}

func init() {
	puzzleLookup["4-1"] = d4p1
	puzzleLookup["4-2"] = d4p2
}

