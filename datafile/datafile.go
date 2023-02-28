// package datafile is used for reading text from files
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// func GetFile is used for reading float-pointing numbers from file
func GetFloat(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

// func GetFile is used for reading strings numbers from file
func GetString(fileName string) ([]string, error) {
	var strs []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			return nil, err
		}
		strs = append(strs, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return strs, nil
}
