// input provides libs around handing input files
package input

import (
	"bufio"
	"os"
	"strconv"
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
	value int
	err   error
}

func StringChanToIntChan(in <-chan string) <-chan IntOrErr {
	out := make(chan IntOrErr)
	go func() {
		defer close(out)

		for line := range in {
			item := IntOrErr{}
			v, err := strconv.ParseUint(line, 10, 64)
			if err != nil {
				item.err = err
			} else {
				item.value = int(v)
			}

			out <- item
		}
	}()

	return out
}
