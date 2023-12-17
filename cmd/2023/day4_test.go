package twentyTwentyThree

import (
	"reflect"
	"testing"
)

func TestNewD4Card(t *testing.T) {
	table := []struct {
		name     string
		line     string
		expected d4Card
	}{
		{"empty", "", d4Card{0, []int64{}, []int64{}}},
		{"id", "Card 4:", d4Card{4, []int64{}, []int64{}}},
		{"strange id", "Card  04  :", d4Card{4, []int64{}, []int64{}}},
		{"winners", "Card 1: 3", d4Card{1, []int64{3}, []int64{}}},
		{"numbers", "Card 5555: 444    | 5", d4Card{5555, []int64{444}, []int64{5}}},
	}

	for _, test := range table {
		actual := NewD4Card(test.line)

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%s failed, got %v but expected %v", test.name, actual, test.expected)
		}
	}
}

func TestD4Card_Score(t *testing.T) {
	table := []struct {
		name     string
		card     d4Card
		expected float64
	}{
		{"empty", d4Card{}, 0},
		{"one non-winner", d4Card{0, []int64{1}, []int64{2}}, 0},
		{"one winner", d4Card{0, []int64{1}, []int64{1}}, 1},
		{"two winner", d4Card{0, []int64{1}, []int64{1, 1}}, 2},
		{"two separate winner", d4Card{0, []int64{1, 2}, []int64{1, 2}}, 2},
		{"three winners", d4Card{0, []int64{1, 2, 3}, []int64{1, 2, 3}}, 4},
	}

	for _, test := range table {
		actual := test.card.Score()

		if actual != test.expected {
			t.Errorf("%s failed, got %v but expected %v", test.name, actual, test.expected)
		}
	}
}

func TestD4Card_Matches(t *testing.T) {
	table := []struct {
		name     string
		card     d4Card
		expected int
	}{
		{"empty", d4Card{}, 0},
		{"one non-matches", d4Card{0, []int64{1}, []int64{2}}, 0},
		{"one matches", d4Card{0, []int64{1}, []int64{1}}, 1},
		{"two matches", d4Card{0, []int64{1}, []int64{1, 1}}, 2},
		{"two separate matches", d4Card{0, []int64{1, 2}, []int64{1, 2}}, 2},
		{"three matches", d4Card{0, []int64{1, 2, 3}, []int64{1, 2, 3}}, 3},
	}

	for _, test := range table {
		actual := test.card.Matches()

		if actual != test.expected {
			t.Errorf("%s failed, got %v but expected %v", test.name, actual, test.expected)
		}
	}
}
