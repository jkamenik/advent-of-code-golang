package twentyTwentyThree

import (
	"testing"
)

func TestD3BoardHasAdjacentSymbol(t *testing.T) {
	table := []struct {
		name     string
		board    d3Board
		x        int
		y        int
		expected bool
	}{
		{"empty", d3Board{}, 0, 0, false},
		{"simple", d3Board{"1"}, 0, 0, false},
		{"left symbol", d3Board{"%1."}, 0, 1, true},
		{"right symbol", d3Board{".1^"}, 0, 1, true},
		{"up left symbol", d3Board{"*..", ".1.", "..."}, 1, 1, true},
		{"up symbol", d3Board{".*.", ".1.", "..."}, 1, 1, true},
		{"up right symbol", d3Board{"..#", ".1.", "..."}, 1, 1, true},
		{"down left symbol", d3Board{"...", ".1.", "#.."}, 1, 1, true},
		{"down symbol", d3Board{"...", ".1.", ".@."}, 1, 1, true},
		{"down right symbol", d3Board{"...", ".1.", "..!"}, 1, 1, true},
	}

	for _, test := range table {
		actual := test.board.hasAdjacentSymbol(test.x, test.y)

		if actual != test.expected {
			t.Errorf("%s failed, got %v but expected %v", test.name, actual, test.expected)
		}
	}
}

func TestSymbol(t *testing.T) {
	table := []struct {
		name     string
		char     byte
		expected bool
	}{
		{"zero", '0', false},
		{"nine", '9', false},
		{"period", '.', false},
		{"star", '*', true},
		{"hash", '#', true},
	}

	for _, test := range table {
		actual := symbol(test.char)

		if actual != test.expected {
			t.Errorf("%s failed, got %v but expected %v", test.name, actual, test.expected)
		}
	}
}
