package datafile

import (
	"bufio"
	"os"
)

// GetStrings function reads a string from each line of a file.
func GetStrings(fileName string) ([]string, error) {
	strings := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return strings, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		string := scanner.Text()
		strings = append(strings, string)
	}
	err = file.Close()
	if err != nil {
		return strings, err
	}
	if scanner.Err() != nil {
		return strings, err
	}
	return strings, nil
}
