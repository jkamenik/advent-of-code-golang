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

func TestD3Number_isNear(t *testing.T) {
	b := d3Board{}
	table := []struct {
		name     string
		num    	d3Number
		x        int
		y        int
		expected bool
	}{
		{"not near - left", d3Number{b,0,2,0,""}, 0, 0, false},
		{"not near - right", d3Number{b,0,0,0,""}, 0, 2, false},
		{"not near - too high", d3Number{b,3,0,0,""}, 0, 0, false},
		{"not near - too low", d3Number{b,0,0,0,""}, 3, 0, false},

		{"not near - up left", d3Number{b,1,5,5,"12345"}, 0, 3, false},
		{"not near - up right", d3Number{b,1,5,5,"12345"}, 0, 12, false},

		{"not near - down left", d3Number{b,1,5,5,"12345"}, 2, 12, false},
		{"not near - down right", d3Number{b,1,5,5,"12345"}, 2, 12, false},

		{"left", d3Number{b,0,1,1,"0"}, 0, 0, true},
		{"right", d3Number{b,0,1,1,"0"}, 0, 2, true},

		{"up", d3Number{b,1,1,1,"0"}, 0, 1, true},
		{"up left", d3Number{b,1,1,1,"0"}, 0, 0, true},
		{"up right", d3Number{b,1,1,1,"0"}, 0, 2, true},

		{"up 2nd digits", d3Number{b,1,10,3,"000"}, 0, 11, true},
		{"up 3rd digits", d3Number{b,1,10,3,"000"}, 0, 12, true},

		{"down", d3Number{b,1,1,1,"0"}, 2, 1, true},
		{"down left", d3Number{b,1,1,1,"0"}, 2, 0, true},
		{"down right", d3Number{b,1,1,1,"0"}, 2, 2, true},

		{"down 2nd digits", d3Number{b,1,10,3,"000"}, 2, 11, true},
		{"down 3rd digits", d3Number{b,1,10,3,"000"}, 2, 12, true},
	}

	for _, test := range table {
		actual := test.num.isNear(test.x, test.y)

		if actual != test.expected {
			t.Errorf("'%s' failed, got %v but expected %v", test.name, actual, test.expected)
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
