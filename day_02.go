package adventofcode2020

import (
	"fmt"
	"regexp"
	"strconv"
)

func dayTwoMatch(input string) (int64, int64, rune, string, error) {
	regex, err := regexp.Compile("^([0-9]+)-([0-9]+) ([^ ]): (.*)$")
	if err != nil {
		return -1, -1, '0', "", err
	}
	result := regex.FindStringSubmatch(input)
	if len(result) < 4 {
		return -1, -1, '0', "", fmt.Errorf("No results for input [%s]: %v",
			input, result)
	}

	lower, err := strconv.ParseInt(result[1], 10, 32)
	if err != nil {
		return -1, -1, '0', "", err
	}
	upper, err := strconv.ParseInt(result[2], 10, 32)
	if err != nil {
		return -1, -1, '0', "", err
	}
	char := []rune(result[3])[0]
	pwd := result[4]
	return lower, upper, char, pwd, nil
}

func dayTwo(input []string) (int, error) {
	ok := 0
	for _, v := range input {
		lower, upper, char, pwd, err := dayTwoMatch(v)
		if err != nil {
			return -1, err
		}
		cnt := int64(0)
		for _, v := range pwd {
			if v == char {
				cnt++
			}
		}
		if lower <= cnt && cnt <= upper {
			ok++
		}
	}
	return ok, nil
}

func dayTwoPart2(input []string) (int, error) {
	ok := 0
	for _, v := range input {
		lower, upper, char, pwd, err := dayTwoMatch(v)
		if err != nil {
			return -1, err
		}

		arr := []rune(pwd)

		if arr[lower-1] == char {
			if arr[upper-1] != char {
				ok++
			}
		} else if arr[upper-1] == char {
			ok++
		}
	}
	return ok, nil
}
