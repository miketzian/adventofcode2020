package adventofcode2020

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readFileAsInts(name string) ([]int, error) {
	// name of a file in this package
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	output := make([]int, 0)
	for scanner.Scan() {
		v, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return nil, err
		}
		output = append(output, int(v))
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return output, nil
}

func readFileAsStrings(name string) ([]string, error) {
	// name of a file in this package
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	output := make([]string, 0)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return output, nil
}
