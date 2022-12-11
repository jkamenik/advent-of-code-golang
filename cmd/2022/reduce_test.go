package twentyTwentyTwo

import (
	"errors"
	"testing"

	"github.com/jkamenik/advent-of-code-golang/input"
	"github.com/rs/zerolog/log"
)

func TestReduce(t *testing.T) {
	table := []struct {
		name string
		in   []input.IntOrErr
		out  []int
	}{
		// Simple cases
		{"empty set", []input.IntOrErr{}, []int{}},
		{"simple 0", []input.IntOrErr{
			{Value: 0},
		}, []int{0}},
		{"simple non-zero", []input.IntOrErr{
			{Value: 1},
		}, []int{1}},

		// Simple errors
		{"empty set error", []input.IntOrErr{
			{Err: errors.New("empty set")},
		}, []int{}},
		{"skip the error", []input.IntOrErr{
			{Err: errors.New("skipped")},
			{Value: 0},
		}, []int{0}},

		// math
		{"two items", []input.IntOrErr{
			{Value: 1},
			{Value: 1},
			{Err: errors.New("separator")},
			{Value: 2},
			{Value: 3},
			{Err: errors.New("skipped")},
		}, []int{2, 5}},
	}

	for _, test := range table {
		log.Info().Msgf("Starting Test %v", test.name)
		items := make([]int, 0)
		c := IntOrErrArrayAsChan(test.in)
		r := reduce(c)

		for i := range r {
			items = append(items, i)
		}

		if len(items) != len(test.out) {
			t.Fatalf("Test %v: Items (%v) not the same length as expected (%v)", test.name, items, test.out)
		}

		if len(items) > 0 {
			// If there are items then check them as
			for i, _ := range test.out {
				if test.out[i] != items[i] {
					t.Errorf("Test %v: Item %v (%v) was not the expected value of %v", test.name, i, items[i], test.out[i])
				}
			}
		}
	}
}

func IntOrErrArrayAsChan(in []input.IntOrErr) <-chan input.IntOrErr {
	out := make(chan input.IntOrErr)

	go func() {
		defer close(out)

		for _, v := range in {
			out <- v
		}
	}()

	return out
}
