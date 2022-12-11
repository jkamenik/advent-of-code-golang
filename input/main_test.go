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
