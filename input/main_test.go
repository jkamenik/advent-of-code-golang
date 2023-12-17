package input

import (
	"reflect"
	"testing"
	"unicode"
)

func TestStreamFile(t *testing.T) {
	collector := make([]string, 0)

	c, err := StreamFile("test/simple.txt", 1)
	if err != nil {
		t.Errorf("Streaming a simple file failed, %v", err)
	}

	for line := range c {
		collector = append(collector, line)
	}

	if len(collector) != 4 {
		t.Errorf("A simple file should have had 4 lines, %v", collector)
	}
}

func TestStringChanToIntChan(t *testing.T) {
	collector := make([]IntOrErr, 0)
	simple, err := StreamFile("test/simple.txt", 1)
	if err != nil {
		t.Errorf("Streaming a simple file failed, %v", err)
	}

	ints := StringChanToIntChan(simple)
	for i := range ints {
		collector = append(collector, i)
	}

	if len(collector) != 4 {
		t.Errorf("A simple file should have had 4 lines, %v", collector)
	}
	for _, v := range collector {
		if v.Err == nil {
			t.Errorf("%v should be an error but wasn't", v)
		}
	}

	collector = make([]IntOrErr, 0)
	numbers, err := StreamFile("test/numbers.txt", 1)
	if err != nil {
		t.Errorf("Streaming a numbers file failed, %v", err)
	}

	ints = StringChanToIntChan(numbers)
	for i := range ints {
		collector = append(collector, i)
	}
	if len(collector) != 7 {
		t.Errorf("A simple file should have had 4 lines, %v", collector)
	}
	for _, v := range collector {
		if v.Err != nil {
			t.Errorf("%v should not have errored but did", v)
		}
	}
}

func TestStringChanToFieldChan(t *testing.T) {
	cases := []struct {
		name  string
		want  [][]string
		delim func(rune) bool
		test  []string
	}{
		{"empty", [][]string{}, unicode.IsSpace, []string{}},
		{"easy", [][]string{{"this"}}, unicode.IsSpace, []string{"this"}},
		{"multiple", [][]string{{"this", "and", "that"}}, unicode.IsSpace, []string{"this and that"}},
		{"multiple delimiters", [][]string{{"this", "and", "that"}}, unicode.IsSpace, []string{" this    and  that"}},
	}

	for _, c := range cases {
		idx := 0
		strings := ChanFromStringSlice(c.test)
		fields := StringChanToFieldChan(strings, c.delim)

		// for each fields
		for field := range fields {
			if idx > len(c.want) {
				t.Fatalf("not enough outputs for %s", c.name)
			}

			if !reflect.DeepEqual(field, c.want[idx]) {
				t.Errorf("%s: expected %v but got %v", c.name, c.want[idx], field)
			}

			idx += 1
		}

		if len(c.test) != idx {
			t.Errorf("%s: expected %d items but found %d", c.name, len(c.test), idx)
		}
	}
}

func ChanFromStringSlice(in []string) <-chan string {
	out := make(chan string, len(in))

	go func() {
		defer close(out)

		for _, line := range in {
			out <- line
		}
	}()

	return out
}
