package twentyTwentyThree

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/rs/zerolog/log"
)

var d1FirstRegex = regexp.MustCompile("([1-9])")
var d1LastRegex = regexp.MustCompile(".*([1-9])")
var d1First2Regex = regexp.MustCompile("([1-9]|zero|one|two|three|four|five|six|seven|eight|nine)")
var d1Last2Regex = regexp.MustCompile(".*([1-9]|zero|one|two|three|four|five|six|seven|eight|nine)")

func d1p1(filename string, file <-chan string) (string, error) {
	return d1Solution(d1FirstRegex, d1LastRegex, file)
}

func d1p2(filename string, file <-chan string) (string, error) {
	return d1Solution(d1First2Regex, d1Last2Regex, file)
}

func d1Solution(firstRx *regexp.Regexp, lastRx *regexp.Regexp, file <-chan string) (string, error) {
	sum := uint64(0)
	for line := range file {
		log.Trace().Msgf("line: %s", line)

		matches := firstRx.FindSubmatch([]byte(line))
		log.Trace().Msgf("First matches: %v", matches)
		first := toDigit(string(matches[1]))

		matches = lastRx.FindSubmatch([]byte(line))
		log.Trace().Msgf("Last matches: %v", matches)
		last := toDigit(string(matches[1]))

		log.Debug().Msgf("%s: %s,%s", line, first, last)

		v, err := strconv.ParseUint(fmt.Sprintf("%s%s", first, last), 10, 64)
		if err != nil {
			return "", err
		}

		sum = sum + v
	}

	return fmt.Sprintf("%v", sum), nil
}

func toDigit(s string) string {
	switch s {
	case "zero":
		return "0"
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"

	default:
		return s
	}
}
