package twentyTwentyThree

import (
	"reflect"
	"testing"
)

func TestNewD4Card(t *testing.T) {
	table := []struct {
		name     	 string
		line     	 string
		expected     d4Card
	}{
		{"empty", "", d4Card{"",[]string{},[]string{}}},
		{"id", "Card 4:", d4Card{"4",[]string{},[]string{}}},
		{"strange id", "Card  04  :",d4Card{"04",[]string{},[]string{}}},
		{"winners","Card 1: 3",d4Card{"1",[]string{"3"},[]string{}}},
		{"numbers","Card 5555: 444    | 5",d4Card{"5555",[]string{"444"},[]string{"5"}}},
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
		expected int
	}{
		{"empty", d4Card{}, 0},
		{"one non-winner", d4Card{"",[]string{"1"},[]string{"2"}}, 0},
		{"one winner", d4Card{"",[]string{"1"},[]string{"1"}}, 1},
		{"two winner", d4Card{"",[]string{"1"},[]string{"1", "1"}}, 2},
		{"two separate winner", d4Card{"",[]string{"1","2"},[]string{"1", "2"}}, 2},
	}

	for _, test := range table {
		actual := test.card.Score()

		if actual != test.expected {
			t.Errorf("%s failed, got %v but expected %v", test.name, actual, test.expected)
		}
	}
}

