package file

import (
	"bufio"
	"os"
)

func LoadData(filename string, callback func(string) error) error {
	fp, err := os.Open(filename)
	defer func() {
		fp.Close()
	}()
	if err != nil {
		return err
	}
	bufferReader := bufio.NewReader(fp)
	for {
		line, err := bufferReader.ReadString('\n')
		if err != nil {
			break
		}
		if err = callback(line); err != nil {
			return err
		}
	}
	return nil
}
