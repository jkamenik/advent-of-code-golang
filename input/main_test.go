package input

import "testing"

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
