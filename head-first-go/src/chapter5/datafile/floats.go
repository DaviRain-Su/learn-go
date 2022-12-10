package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([3]float64, error) {
	var number [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return number, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return number, err
		}
		i += 1
	}
	err = file.Close()
	if err != nil {
		return number, err
	}
	if scanner.Err() != nil {
		return number, err
	}
	return number, nil
}
