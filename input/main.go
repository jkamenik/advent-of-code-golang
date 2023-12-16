// input provides libs around handing input files
package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type LineHandler func(string) error

func ScanFile(fileName string, handler LineHandler) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		err = handler(line)
		if err != nil {
			return err
		}
	}

	return nil
}

func StreamFile(fileName string, buffer int) (<-chan string, error) {
	out := make(chan string, buffer)
	file, err := os.Open(fileName)
	if err != nil {
		return out, err
	}

	go func() {
		defer file.Close()
		defer close(out)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			out <- line
		}
	}()

	return out, err
}

// IntOrErr is either the int or an error generated from parsing
type IntOrErr struct {
	Value int
	Err   error
}

func StringChanToIntChan(in <-chan string) <-chan IntOrErr {
	out := make(chan IntOrErr)
	go func() {
		defer close(out)

		for line := range in {
			item := IntOrErr{}
			v, err := strconv.ParseUint(line, 10, 64)
			if err != nil {
				item.Err = err
			} else {
				item.Value = int(v)
			}

			out <- item
		}
	}()

	return out
}


// StringChanToFieldChan converts a string change into a channel of fields
// The isDelimiter function should return true if the character is a delimiter.
// Delimiters are not copied as fields.
func StringChanToFieldChan(in <-chan string, isDelimiter func(rune) bool) <-chan []string {
	out := make(chan []string)

	go func() {
		defer close(out)

		for line := range in {
			fields := strings.FieldsFunc(line, isDelimiter)

			out <- fields
		}
	}()


	return out
}

func FieldsAsInts(fields []string) ([]int64, error) {
	out := make([]int64, len(fields))

	for idx, i := range fields {
		v, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			return out, err
		}

		out[idx] = v
	}

	return out, nil
}
