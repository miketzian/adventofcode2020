package adventofcode2020

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func reverseString(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}
	// return the reversed string.
	return string(rns)
}
func reverseStringArray(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		arr[i], arr[j] = arr[j], arr[i]
	}
}

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

func readFileAsInt64s(name string) ([]int64, error) {
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

	output := make([]int64, 0)
	for scanner.Scan() {
		v, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		output = append(output, v)
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
