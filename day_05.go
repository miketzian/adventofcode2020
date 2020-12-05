package adventofcode2020

import (
	"fmt"
)

func day5DetermineRow(input string) (int, error) {

	if len(input) < 10 {
		return -1, fmt.Errorf("Invalid Input [%s]", input)
	}

	min := 0
	max := 128

	for _, v := range []rune(input[:6]) {
		switch v {
		case 'F':
			max = max - ((max - min) / 2)
		case 'B':
			min = max - ((max - min) / 2)
		default:
			return -1, fmt.Errorf("Invalid Char [%s] for input [%s]",
				string(v), input)
		}
	}
	last := input[6:7]
	if last == "F" {
		return min, nil
	} else if last == "B" {
		return max - 1, nil
	}
	return -1, fmt.Errorf("Invalid Char [%s] for input [%s]",
		last, input)

}

func day5DetermineSeat(input string) (int, error) {

	if len(input) < 10 {
		return -1, fmt.Errorf("Invalid Input [%s]", input)
	}
	min := 0
	max := 8
	for _, v := range []rune(input[7:9]) {
		switch v {
		case 'R':
			min = max - ((max - min) / 2)
		case 'L':
			max = max - ((max - min) / 2)
		default:
			return -1, fmt.Errorf("Invalid Char [%s] for input [%s]",
				string(v), input)
		}
	}
	last := input[9:]
	if last == "L" {
		return min, nil
	} else if last == "R" {
		return max - 1, nil
	}
	return -1, fmt.Errorf("Invalid Char [%s] for input [%s]",
		last, input)
}

func day5(input string) (int, int, int, error) {

	row, err := day5DetermineRow(input)
	if err != nil {
		return -1, -1, -1, err
	}
	col, err := day5DetermineSeat(input)
	if err != nil {
		return -1, -1, -1, err
	}
	return row, col, (row * 8) + col, nil
}
