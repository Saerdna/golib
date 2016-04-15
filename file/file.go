package file

import (
	"bufio"
	"io/ioutil"
	"os"
)

func LoadData(filename string, callback func(string) error) error {
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fp.Close()
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

func LoadDataAll(filename string, callback func([]byte) error) error {
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fp.Close()
	b, err := ioutil.ReadAll(fp)
	if err != nil {
		return err
	}
	err = callback(b)
	return err
}
