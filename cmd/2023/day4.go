package twentyTwentyThree

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

func d4p1(filename string, file <-chan string) (string, error) {
	sum := float64(0)

	for line := range file {
		card := NewD4Card(line)

		sum = sum + card.Score()
	}

	return fmt.Sprintf("%f", sum), nil
}

func d4p2(filename string, file <-chan string) (string, error) {
	cards := []d4Card{}

	for line := range file {
		cards = append(cards, NewD4Card(line))
	}

	counts := make([]int, len(cards)+1)

	for _, card := range cards {
		counts[card.ID] += 1
		matches := card.Matches()
		factor := counts[card.ID]

		log.Trace().Int("matches", matches).Int("factor", factor).Msgf("counts %v", counts)

		for i := card.ID + 1; i <= card.ID+int64(matches); i++ {
			counts[i] += factor
			log.Trace().Str("counts", fmt.Sprintf("%v", counts)).Msgf("incrementing %d by %d", i, factor)
		}
	}

	sum := 0
	for _, count := range counts {
		sum += count
	}

	return fmt.Sprintf("%d", sum), nil
}

type d4Card struct {
	ID      int64
	Winners []int64
	Numbers []int64
}

func (c d4Card) String() string {
	return fmt.Sprintf("{id: '%v', w: %v, n: %v}", c.ID, c.Winners, c.Numbers)
}

func NewD4Card(line string) (card d4Card) {
	card.Winners = []int64{}
	card.Numbers = []int64{}

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

func DigitsFromString(s string) []int64 {
	rtn := []int64{}

	acc := ""

	for _, r := range s {
		log.Trace().Str("acc", acc).Str("digits", fmt.Sprintf("%v", rtn)).Msgf("rune: '%v'", string(r))

		if r >= '0' && r <= '9' {
			acc = acc + string(r)
		} else {
			log.Trace().Str("acc", acc).Msg("Not a number, flush accumulator")
			if acc != "" {
				n, err := strconv.ParseInt(acc, 10, 64)
				if err != nil {
					panic(err)
				}

				rtn = append(rtn, n)
				acc = ""
			}
		}
	}

	if acc != "" {
		n, err := strconv.ParseInt(acc, 10, 64)
		if err != nil {
			panic(err)
		}

		rtn = append(rtn, n)
	}

	log.Trace().Msgf("digits %v", rtn)

	return rtn
}

func (card d4Card) Score() float64 {
	matches := card.Matches()
	log.Trace().Msgf("matches %d", matches)
	if matches == 0 {
		return 0
	} else {
		return math.Pow(2, float64(matches-1))
	}
}

func (card d4Card) Matches() int {
	matches := 0

	for _, winner := range card.Winners {
		for _, have := range card.Numbers {
			if winner == have {
				log.Trace().Msgf("Found match %v", winner)
				matches = matches + 1
			}
		}
	}

	return matches
}
