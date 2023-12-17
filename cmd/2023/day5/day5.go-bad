// Below was a naive implementation expanding all the
package day5

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/jkamenik/advent-of-code-golang/input"
	"github.com/rs/zerolog/log"
)

func Part1(filename string, file <-chan string) (string, error) {
	seeds := newSeedMap()
	state := stateSeed

	for fields := range input.StringChanToFieldChan(file, unicode.IsSpace) {
		if state == nil {
			log.Info().Strs("fields", fields).Msg("unknown state")
			break
		}

		state = state(seeds, fields)
	}

	log.Debug().Stringer("map", seeds).Msg("Seed Map")

	return fmt.Sprintf("min: %v", seeds.lowestLocation()), nil
}

func Part2(filename string, file <-chan string) (string, error) {
	return "", errors.New("not implemented")
}

type seedMap struct {
	seeds []int64

  seedToSoil toMap
  soilToFertilizer toMap
  fertilizerToWater toMap
  WaterToLight toMap
  lightToTemperature toMap
  temperatureToHumidity toMap
  humidityToLocation toMap
}

func newSeedMap() *seedMap {
	seeds := seedMap{
		seeds: []int64{},

    seedToSoil: toMap{},
    soilToFertilizer: toMap{},
    fertilizerToWater: toMap{},
    WaterToLight: toMap{},
    lightToTemperature: toMap{},
    temperatureToHumidity: toMap{},
    humidityToLocation: toMap{},
	}

	return &seeds
}

func (s *seedMap) String() string {
	return fmt.Sprintf("{seeds: %v, seedToSoil: %v, soilToFertilizer: %v, fertilizerToWater: %v, waterToLight: %v, lightToTemperature: %v, temperatureToHumidity: %v, humidityToLocation: %v}", s.seeds, s.seedsToSoil, s.soilToFertilizer, s.fertilizerToWater, s.waterToLight, s.lightToTemperature, s.temperatureToHumidity, s.humidityToLocation)
}


type toMap []entry
type entry struct {
  src int64
  dest int64
  count int64
}

// This defined a finite state machine
type stateFn func(*seedMap, []string) stateFn

func (s seedMap) lowestLocation() int64 {
	minLoc := int64(0)

	return minLoc
}

func stateSeed(seeds *seedMap, fields []string) stateFn {
	log.Trace().Strs("Seeds",fields).Msg("seeds")

	ints, err := input.FieldsAsInts(fields[1:])
	if err != nil {
		log.Error().Err(err).Msg("unable to convert fields to int")
		return nil
	}

	seeds.seeds = ints

  // now move on to the transitions
	return stateMap
}

func stateMap(seeds *seedMap, fields []string) stateFn {
	log.Trace().Strs("fields",fields).Msg("stateMap")

	if len(fields) < 1 {
		// empty line
		return stateMap
	}

	switch fields[0] {
	case "seed-to-soil":
		return stateSeedToSoil
	case "soil-to-fertilizer":
		return stateSoilToFertilizer
	case "fertilizer-to-water":
		return stateFertilizerToWater
	case "water-to-light":
		return stateWaterToLight
	case "light-to-temperature":
		return stateLightToTemperature
	case "temperature-to-humidity":
		return stateTemperatureToHumidity
	case "humidity-to-location":
		return stateHumidityToLocation
	default:
		log.Trace().Str("field", fields[0]).Msg("No state match found")
		return stateMap
	}
}

func updateState(place map[int64]int64, fields []string) error {
	log.Trace().Strs("fields",fields).Msg("updateState")

	if len(fields) != 3 {
		// give statemap a chance
		log.Warn().Strs("fields", fields).Msg("invalid line")
		return nil
	}

	ints, err := input.FieldsAsInts(fields)
	if err != nil {
		log.Error().Err(err).Msg("unable to convert fields to int")
		return err
	}

	val := ints[0]
	idx := ints[1]
	count := ints[2]

	for i := int64(0); i < count; i++ {
		log.Trace().Int64("idx", idx).Int64("val", val).Msg("setting map")
		place[idx] = val

		val += 1
		idx += 1
	}

	return nil
}

func stateSeedToSoil(seeds *seedMap, fields []string) stateFn {
	log.Trace().Strs("fields",fields).Msg("stateSeedToSoil")

	if len(fields) != 3 {
		// give statemap a chance
		log.Warn().Strs("fields", fields).Msg("invalid line")
		return stateMap
	}

	err := updateState(seeds.seedsToSoil, fields)
	if err != nil {
		log.Error().Err(err).Msg("unable to update state")
		return nil
	}

	return stateSeedToSoil
}

func stateSoilToFertilizer(seeds *seedMap, fields []string) stateFn {
	log.Trace().Strs("fields",fields).Msg("stateSoilToFertilizer")

	if len(fields) != 3 {
		// give statemap a chance
		log.Warn().Strs("fields", fields).Msg("invalid line")
		return stateMap
	}

	err := updateState(seeds.soilToFertilizer, fields)
	if err != nil {
		log.Error().Err(err).Msg("unable to update state")
		return nil
	}

	return stateSoilToFertilizer
}

func stateFertilizerToWater(seeds *seedMap, fields []string) stateFn {
	log.Trace().Strs("fields",fields).Msg("stateFertilizerToWater")

	if len(fields) != 3 {
		// give statemap a chance
		log.Warn().Strs("fields", fields).Msg("invalid line")
		return stateMap
	}

	err := updateState(seeds.fertilizerToWater, fields)
	if err != nil {
		log.Error().Err(err).Msg("unable to update state")
		return nil
	}

	return stateFertilizerToWater
}

func stateWaterToLight(seeds *seedMap, fields []string) stateFn {
	log.Trace().Strs("fields",fields).Msg("stateWaterToLight")

	if len(fields) != 3 {
		// give statemap a chance
		log.Warn().Strs("fields", fields).Msg("invalid line")
		return stateMap
	}

	err := updateState(seeds.waterToLight, fields)
	if err != nil {
		log.Error().Err(err).Msg("unable to update state")
		return nil
	}

	return stateWaterToLight
}

func stateLightToTemperature(seeds *seedMap, fields []string) stateFn {
	log.Trace().Strs("fields",fields).Msg("stateLightToTemperature")

	if len(fields) != 3 {
		// give statemap a chance
		log.Warn().Strs("fields", fields).Msg("invalid line")
		return stateMap
	}

	err := updateState(seeds.lightToTemperature, fields)
	if err != nil {
		log.Error().Err(err).Msg("unable to update state")
		return nil
	}

	return stateLightToTemperature
}

func stateTemperatureToHumidity(seeds *seedMap, fields []string) stateFn {
	log.Trace().Strs("fields",fields).Msg("stateTemperatureToHumidity")

	if len(fields) != 3 {
		// give statemap a chance
		log.Warn().Strs("fields", fields).Msg("invalid line")
		return stateMap
	}

	err := updateState(seeds.temperatureToHumidity, fields)
	if err != nil {
		log.Error().Err(err).Msg("unable to update state")
		return nil
	}

	return stateTemperatureToHumidity
}

func stateHumidityToLocation(seeds *seedMap, fields []string) stateFn {
	log.Trace().Strs("fields",fields).Msg("stateHumidityToLocation")

	if len(fields) != 3 {
		// give statemap a chance
		log.Warn().Strs("fields", fields).Msg("invalid line")
		return stateMap
	}

	err := updateState(seeds.humidityToLocation, fields)
	if err != nil {
		log.Error().Err(err).Msg("unable to update state")
		return nil
	}

	return stateHumidityToLocation
}
