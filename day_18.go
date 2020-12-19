package adventofcode2020

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day18Calc(input string) (int, error) {
	val := 0
	nextOp := "+"
	ix := 0
	lastIx := 0
	for ix < len(input) {
		switch input[ix : ix+1] {
		case "+":
			fallthrough
		case "*":
			nextVal, err := strconv.Atoi(input[lastIx:ix])
			if err != nil {
				return 0, err
			}
			switch nextOp {
			case "+":
				val += nextVal
			case "*":
				val *= nextVal
			}
			nextOp = input[ix : ix+1]
			lastIx = ix + 1
		}
		ix++
	}
	nextVal, err := strconv.Atoi(input[lastIx:])
	if err != nil {
		return 0, err
	}
	switch nextOp {
	case "+":
		val += nextVal
	case "*":
		val *= nextVal
	}
	// fmt.Printf("input [%s]=%d\n", input, val)
	return val, nil
}

func day18CalcPart2(input string) (int, error) {

	// the addition operations must be done first in this round
	// there are no parenthesis or spaces at this point

	addRegex, err := regexp.Compile(`([0-9]+)\+([0-9]+)`)
	if err != nil {
		return 0, err
	}
	m := addRegex.FindStringSubmatchIndex(input)
	for len(m) > 0 {

		num, err := day18Calc(input[m[0]:m[1]])
		if err != nil {
			return 0, err
		}
		input = fmt.Sprintf("%s%d%s", input[0:m[0]], num, input[m[1]:])
		m = addRegex.FindStringSubmatchIndex(input)
	}
	// at this point it will be all multiplication or just a number
	return day18Calc(input)
}

func day18Line(input string, part2 bool) (int, error) {

	var calcFn func(input string) (int, error)
	if part2 {
		calcFn = day18CalcPart2
	} else {
		calcFn = day18Calc
	}

	// first, factor out the parenthesis
	type Group struct {
		data string
		head *Group
	}
	group := &Group{"", nil}
	for _, x := range strings.Replace(input, " ", "", -1) {
		switch x {
		case '(':
			group = &Group{"", group}
		case ')':
			val, err := calcFn(group.data)
			if err != nil {
				return 0, err
			}
			if group.head == nil {
				return 0, fmt.Errorf("mismatched parenthesis too many ')'")
			}
			group = group.head
			group.data += fmt.Sprintf("%d", val)
		default:
			group.data += string(x)
		}
	}
	if group.head != nil {
		return 0, fmt.Errorf("mismatched parenthesis too many '('")
	}
	nextVal, err := calcFn(group.data)
	if err != nil {
		return 0, err
	}
	return nextVal, nil
}

func day18(input []string, part2 bool) (int, error) {
	total := 0
	for _, line := range input {
		line, err := day18Line(line, part2)
		if err != nil {
			return 0, err
		}
		total += line
	}
	return total, nil
}
