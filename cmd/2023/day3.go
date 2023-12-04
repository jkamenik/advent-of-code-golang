package twentyTwentyThree

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

var (
	NoNumberErr = errors.New("No Number Found")
)

func d3p1(filename string, file <-chan string) (string, error) {
	board := newD3Board(file)
	log.Trace().Msgf("%v", board)

	sum := int64(0)
	for num := range board.iterate() {
		log.Debug().Msgf("Number: %v", num)

		if num.isSymbolAdjacent() {
			n, err := strconv.ParseInt(num.Number, 10, 64)
			if err != nil {
				return "", err
			}
			sum = sum + n
		} else {
			log.Info().Msgf("No symbol near %v", num)
		}
	}

	return fmt.Sprintf("%d", sum), nil
}

func d3p2(filename string, file <-chan string) (string, error) {
	board := newD3Board(file)
	log.Trace().Msgf("%v", board)

	numbers := []d3Number{}
	for n := range board.iterate() {
		numbers = append(numbers, n)
	}

	sum := int64(0)
	for gear := range board.gears() {
		log.Debug().Msgf("Gear: %v", gear)

		ratio := []d3Number{}

		// see if there are any numbers touching it
		for _, num := range numbers {
			// bail early
			if len(ratio) == 2 {
				break
			}

			if num.isNear(gear.X, gear.Y) {
				ratio = append(ratio, num)
			}
		}

		log.Debug().Msgf("Ratio: %v", ratio)
		if len(ratio) < 2 {
			log.Info().Msgf("Gear %v doesn't have a ratio", gear)
			continue
		}

		n1, err := strconv.ParseInt(ratio[0].Number, 10, 64)
		if err != nil {
			return "", err
		}

		n2, err := strconv.ParseInt(ratio[1].Number, 10, 64)
		if err != nil {
			return "", err
		}

		sum = sum + (n1 * n2)
	}

	return fmt.Sprintf("%d", sum), nil
}

type d3Board []string

func newD3Board(lines <-chan string) d3Board {
	board := d3Board{}

	for line := range lines {
		board = append(board, line)
	}

	return board
}

func (b d3Board) iterate() <-chan d3Number {
	rtn := make(chan d3Number)

	go func() {
		for x := 0; x < len(b); x++ {
			line := b[x]
			num := newD3Number(b)
			for y := 0; y < len(line); y++ {
				r := line[y]
				// log.Trace().Msgf("%d, %d: %v", x, y, string(line[y]))

				if r >= '0' && r <= '9' {
					num.Update(x, y, r)
					log.Trace().Msgf("%v", num)
				} else {
					// flush the last number if valid and start over
					if num.X != -1 {
						rtn <- num
					}
					num = newD3Number(b)
				}
			}

			// if we still have a number then flush it
			if num.X != -1 {
				rtn <- num
			}
		}

		close(rtn)
	}()

	return rtn
}

func (b d3Board) gears() <-chan d3Number {
	rtn := make(chan d3Number)

	go func() {
		for x := 0; x < len(b); x++ {
			line := b[x]
			num := newD3Number(b)
			for y := 0; y < len(line); y++ {
				r := line[y]
				// log.Trace().Msgf("%d, %d: %v", x, y, string(line[y]))

				if r == '*' {
					num.Update(x, y, r)
					log.Trace().Msgf("%v", num)
				} else {
					// flush the last number if valid and start over
					if num.X != -1 {
						rtn <- num
					}
					num = newD3Number(b)
				}
			}

			// if we still have a number then flush it
			if num.X != -1 {
				rtn <- num
			}
		}

		close(rtn)
	}()

	return rtn
}

func (b d3Board) hasAdjacentSymbol(x, y int) bool {
	log.Trace().Msgf("Checking symbols around %d, %d", x, y)

	up := x - 1
	down := x + 1
	left := y - 1
	right := y + 1

	// empty board or row
	if len(b) <= 0 {
		log.Trace().Msg("Empty board")
		return false
	} else if len(b[x]) <= 0 {
		log.Trace().Msg("Empty row")
		return false
	}

	// Check around myself
	if left >= 0 && symbol(b[x][left]) {
		return true
	}
	if right < len(b[x]) && symbol(b[x][right]) {
		return true
	}

	if up >= 0 {
		if left >= 0 && symbol(b[up][left]) {
			return true
		}

		if symbol(b[up][y]) {
			return true
		}

		if right < len(b[up]) && symbol(b[up][right]) {
			return true
		}
	}

	if down < len(b) {
		if left >= 0 && symbol(b[down][left]) {
			return true
		}

		if symbol(b[down][y]) {
			return true
		}

		if right < len(b[down]) && symbol(b[down][right]) {
			return true
		}
	}

	return false
}

func symbol(b byte) bool {
	if b >= '0' && b <= '9' || b == '.' {
		return false
	}

	return true
}

type d3Number struct {
	Board  d3Board
	X      int
	Y      int
	Length int
	Number string
}

func newD3Number(board d3Board) (num d3Number) {
	num.Board = board
	num.X = -1
	num.Y = -1
	num.Length = 0

	return
}

func (n d3Number) String() string {
	return fmt.Sprintf("{%d, %d: %s (%d)}", n.X, n.Y, n.Number, n.Length)
}

func (n *d3Number) Update(x, y int, c byte) {
	if n.X == -1 {
		n.X = x
	}

	if n.Y == -1 {
		n.Y = y
	}

	n.Length = n.Length + 1
	n.Number = n.Number + string(c)
}

func (n d3Number) isSymbolAdjacent() bool {
	// for each character check all adjacent cells
	for y := 0; y < n.Length; y++ {
		if n.Board.hasAdjacentSymbol(n.X, n.Y+y) {
			return true
		}
	}

	return false
}

func (n d3Number) isNear(x, y int) bool {
	// same row
	if x == n.X {
		if y == n.Y - 1  || y == n.Y + n.Length {
			return true
		}
	}

	// above
	if x == n.X - 1{
		if y >= n.Y - 1  && y <= n.Y + n.Length{
			return true
		}
	}

	// below
	if x == n.X + 1 {
		if y >= n.Y - 1  && y <= n.Y + n.Length{
			return true
		}
	}

	return false
}

func init() {
	puzzleLookup["3-1"] = d3p1
	puzzleLookup["3-2"] = d3p2
}
