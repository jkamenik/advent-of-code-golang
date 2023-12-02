package twentyTwentyThree

import (
	"fmt"
	"strconv"
	"regexp"

	"github.com/rs/zerolog/log"
)

var d1Regex = regexp.MustCompile("[1-9]")
var d1_2Regex = regexp.MustCompile("[1-9]|zero|one|two|three|four|five|six|seven|eight|nine")

func d1p1(filename string, file <-chan string)(string, error) {
	return d1Solution(d1Regex, file)
}

func d1p2(filename string, file <-chan string)(string, error) {
	return d1Solution(d1_2Regex, file)
}

func d1Solution(regex *regexp.Regexp, file <- chan string)(string, error) {
	sum := uint64(0)
	for line := range file {
		log.Trace().Msgf("line: %s", line)
		first := ""
		last := ""

		all := regex.FindAllString(line, -1)

		for _, match := range all {
			log.Trace().Msgf("number: %v", match)
			last = toDigit(match)
			if first == "" {
				first = last
			}
		}

		log.Debug().Msgf("%s: %s%s",line, first,last)

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


func init() {
	puzzleLookup["1-1"] = d1p1
	puzzleLookup["1-2"] = d1p2
}
