package twentyTwentyTwo

import (
	"github.com/jkamenik/advent-of-code-golang/input"
	"github.com/rs/zerolog/log"
)

// reduce takes a channel of IntOrErr objects
// accumulates them until there is an error or
// the input channel is closed.  It emits the
// accumulator.
func reduce(in <-chan input.IntOrErr) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		hasOne := false
		accumulator := 0

		for i := range in {
			log.Info().Err(i.Err).Int("value", i.Value).Msg("received input")

			if i.Err == nil {
				hasOne = true
				accumulator = accumulator + i.Value

				log.Debug().Int("accumulator", accumulator).Msg("Not an error, accumulate")
			} else {
				log.Info().Err(i.Err).Msg("value had an error, emitting current")
				if hasOne {
					out <- accumulator
				}

				hasOne = false
				accumulator = 0
			}
		}

		// have to emit one last one, if there is one
		if hasOne {
			out <- accumulator
		}
	}()

	return out
}
