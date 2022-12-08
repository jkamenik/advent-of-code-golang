package twentyTwentyTwo

import "testing"

func TestNewElf(t *testing.T) {
	e := NewElf()

	if e.calories == nil {
		t.Fatalf("Elf %v has a nil array", e)
	}

	if len(e.calories) != 0 {
		t.Errorf("Elf %v was not created empty", e)
	}
}

func TestElfAddCalories(t *testing.T) {
	e := NewElf()

	err := e.AddCalories("1")
	if err != nil {
		t.Errorf("1 should have been a valid calorie: %v", err)
	}
	if e.calories[0] != 1 {
		t.Errorf("Calorie should have been 1, but was %v", e.calories[0])
	}

	err = e.AddCalories("a")
	if err == nil {
		t.Errorf("a should not have been a valid calorie")
	}

	err = e.AddCalories("2")
	if len(e.calories) != 2 {
		t.Errorf("2nd calorie wasn't added, %v", e)
	}
}

func TestNewElves(t *testing.T) {
	e := NewElves()

	if e.all == nil {
		t.Fatalf("Elves %v has a nil array", e)
	}

	if len(e.all) != 0 {
		t.Errorf("Elves %v was not created empty", e)
	}
}

func TestElvesReadLines(t *testing.T) {
	e := NewElves()
	lines := []string{""}
	for _, line := range lines {
		e.ReadLine(line)
	}
	if len(e.all) != 1 {
		t.Fatalf("there should have been 1 elf, %v", e)
	}

	e = NewElves()
	lines = []string{"", ""}
	for _, line := range lines {
		e.ReadLine(line)
	}
	if len(e.all) != 2 {
		t.Fatalf("there should have been 2 elves, %v", e)
	}

	e = NewElves()
	lines = []string{
		"100",
		"200",
	}
	for _, line := range lines {
		e.ReadLine(line)
	}
	if len(e.all) != 1 {
		t.Fatalf("there should have been 1 elf, %v", e)
	}
	if len(e.all[0].calories) != 2 {
		t.Errorf("The number of calories for elf 1 should have been 2, %v", e)
	}
	total := e.all[0].Total()
	if total != 300 {
		t.Errorf("The total calories for elf 1 should have been 300 was %v", total)
	}
}
