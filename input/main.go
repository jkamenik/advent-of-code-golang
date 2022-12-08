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

		err = handle(line)
		if err != nil {
			return err
		}
	}

	return nil
}