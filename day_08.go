package adventofcode2020

import (
	"fmt"
	"regexp"
	"strconv"
)

// Day8 ...
type Day8 struct {
	lineParser *regexp.Regexp
}

func (o *Day8) parseLine(input string) (string, int, error) {
	fmt.Printf("Input Parsing [%s]\n", input)
	result := o.lineParser.FindStringSubmatch(input)
	if len(result) < 3 {
		return "", -1, fmt.Errorf("Input failed to parse [%s]", input)
	}
	num64, err := strconv.ParseInt(result[3], 10, 32)
	if err != nil {
		return "", -1, err
	}
	if result[2] == "-" {
		num64 *= -1
	}
	return result[1], int(num64), nil
}

func (o *Day8) runLoop(input []string, errOnSeen bool) (int, error) {

	seen := make(map[int]bool)
	acc := 0
	ix := 0

	// The program is supposed to terminate by attempting to execute
	// an instruction immediately after the last instruction in the file.
	for ix != len(input) {
		if _, prs := seen[ix]; prs {
			// in part one, finding the loop is the success condition
			if errOnSeen {
				return -1, fmt.Errorf("Line already seen: [%s]", input[ix])
			}
			return acc, nil
		}
		if len(input) < ix {
			return -1, fmt.Errorf("No input for input [%d]", ix)
		}
		op, num, err := o.parseLine(input[ix])
		if err != nil {
			return -1, err
		}
		seen[ix] = true
		switch {
		case op == "nop":
			ix++
		case op == "acc":
			acc += num
			ix++
		case op == "jmp":
			ix += num
			fmt.Printf("Ix is now [%d]\n", ix)
		}
	}
	return acc, nil
}

func dayEight(input []string) (int, error) {

	obj := Day8{lineParser: regexp.MustCompile(`^(nop|jmp|acc) (\+|\-)([0-9]+)$`)}

	//     day_08_test.go:40: Result: 1610
	num, err := obj.runLoop(input, false)
	if err != nil {
		return -1, err
	}
	return num, err
}

func dayEightPart2(input []string) (int, error) {
	swap := 0
	obj := Day8{lineParser: regexp.MustCompile(`^(nop|jmp|acc) (\+|\-)([0-9]+)$`)}

	var oldLine string
	for swap < len(input) {
		oldLine = input[swap]
		if oldLine[0:3] == "nop" {
			// try swapping the command
			input[swap] = "jmp" + oldLine[3:]
			res, err := obj.runLoop(input, true)
			if err == nil {
				return res, nil
			}
			// otherwise, put it back
			input[swap] = oldLine

		} else if oldLine[0:3] == "jmp" {
			// try swapping the command
			input[swap] = "nop" + oldLine[3:]
			res, err := obj.runLoop(input, true)
			if err == nil {
				return res, nil
			}
			// otherwise, put it back
			input[swap] = oldLine
		}
		// and loop
		swap++
	}
	return -1, fmt.Errorf("no match found")
}
