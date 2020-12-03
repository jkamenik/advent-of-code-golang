package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	clearMarker = "."[0]
	treeMarker  = "#"[0]
	hitMarker   = "X"[0]
	missMarker  = "O"[0]
)

// field represents a field of trees in the form of "." for emtpy space and
// "#" for a tree.
type field struct {
	field       []string
	walkedField []string
	right       int
	down        int

	misses int
	hits   int
	steps  int
}

func newField(fileName string) (*field, error) {
	lines := make([]string, 0)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return &field{field: lines}, nil
}

// Using the walk values it guesses at the required field width.
// Example:
//   Field is 5 wide and 4 tall.
//   Right is 4 and down is 2
//   Then we move 8 to the right
//		(4 tall / 2 down) * 4 right
//   Since the field is 5 wide then we double the field field width
//		(i.e, ceiling(8 / 5))
func (f *field) resize() {
	height := len(f.field)
	if height == 0 {
		// cannot walk an empty field
		f.walkedField = make([]string, 0)
		return
	}

	width := len(f.field[0])
	maxRightMovement := (height / f.down) * f.right
	fieldReplicationCount := int(math.Ceil(float64(maxRightMovement) / float64(width)))

	fmt.Printf("height (%d) x width (%d)\nmax right place (%d)\nwidth replication (%d)\n", height, width, maxRightMovement, fieldReplicationCount)

	f.walkedField = make([]string, height)

	for idx, line := range f.field {
		f.walkedField[idx] = strings.Repeat(line, fieldReplicationCount)
	}
}

// Walks the trail, modifing the field to mark the hits and misses.
// It returns the total number of steps taken, the hits and the misses.
func (f *field) walk(right, down int) (steps, hits, misses int, err error) {
	if right <= 0 || down <= 0 {
		return 0, 0, 0, fmt.Errorf("Either right or down value out of range")
	}

	if len(f.field)%down != 0 {
		return 0, 0, 0, fmt.Errorf("Down (%d) is not a multiple of the field height (%d)", down, len(f.field))
	}

	f.hits = 0
	f.misses = 0
	f.steps = 0
	f.right = right
	f.down = down
	f.resize()

	fmt.Printf("clear (%v), tree (%v), miss (%v), hit (%v) markers\n", clearMarker, treeMarker, missMarker, hitMarker)

	for i := 0; i < len(f.walkedField); i += down {
		f.steps++

		j := i * right

		row := f.walkedField[i]
		character := row[j]

		if character == treeMarker {
			f.hits++
			f.walkedField[i] = replaceRune(row, hitMarker, j)
		} else {
			f.misses++
			f.walkedField[i] = replaceRune(row, missMarker, j)
			// row[j] = missMarker
		}

		fmt.Printf("%d x %d (%v)\n", i, j, character)
	}

	return f.steps, f.hits, f.misses, nil
}

// String renders the field as new line separated
func (f *field) String() string {
	if len(f.walkedField) <= 0 {
		return fmt.Sprintf("%s\n(unwalked)", strings.Join(f.field, "\n"))
	}

	return strings.Join(f.walkedField, "\n")
}

func replaceRune(input string, character byte, index int) string {
	runes := make([]byte, len(input))
	for idx, c := range input {
		if idx == index {
			runes = append(runes, character)
		} else {
			runes = append(runes, byte(c))
		}
	}

	return string(runes)
}