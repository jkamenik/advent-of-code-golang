// input provides libs around handing input files
package input

import (
	"bufio"
	"os"
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
