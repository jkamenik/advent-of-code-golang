package day5

import (
	"errors"
	"strings"

	"github.com/rs/zerolog/log"
)

func Part1(filename string, file <-chan string) (string, error) {
	seeds := newSeedMap()
	state := stateSeed

	for line := range file {
		if state == nil {
			log.Info().Str("line", line).Msg("unknown state")
			break
		}

		state = state(seeds, line)
	}

	return "done", nil
}

func Part2(filename string, file <-chan string) (string, error) {
	return "", errors.New("not implemented")
}

type seedMap struct {
	seedsToSoil map[string]string
}
type stateFn func(*seedMap, string) stateFn

func newSeedMap() *seedMap {
	seeds := seedMap{}

	return &seeds
}

func stateSeed(seeds *seedMap, line string) stateFn {
	// line is "seeds: " + fields of seeds
	f := strings.Fields(line)

	log.Trace().Strs("Seeds",f).Msg("seeds")

	return stateMap
}

func stateMap(seeds *seedMap, line string) stateFn {
	f := strings.Fields(line)
	log.Trace().Strs("fields",f).Msg("stateMap")

	if len(f) < 1 {
		// empty line
		return stateMap
	}

	switch f[0] {
	case "seed-to-soil":
		return stateSeedToSoil
	default:
		log.Trace().Str("field", f[0]).Msg("No state match found")
		return stateMap
	}
}

func stateSeedToSoil(seeds *seedMap, line string) stateFn {
	f := strings.Fields(line)
	log.Trace().Strs("fields",f).Msg("stateSeedToSoil")

	if len(f) != 3 {
		// give statemap a chance
		return stateMap(seeds, line)
	}

	soil := f[0]
	seed := f[1]
	count := f[2]

	// assume I am still correct
	return stateSeedToSoil
}
