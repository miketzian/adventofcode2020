package adventofcode2020

import (
	"strconv"
	"strings"
)

func day15(startNums string, loops int) (int, error) {

	history := make(map[int][]int)

	loop := 0
	lastNum := 0
	for _, v := range strings.Split(startNums, ",") {
		num, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		history[num] = []int{loop, -1}
		lastNum = num
		loop++
	}

	for loop < loops {
		thisNum := 0
		// lastNum is always in history
		if history[lastNum][1] != -1 {
			// not the first time it was spoken
			thisNum = history[lastNum][1] - history[lastNum][0]
		}
		if _, prs := history[thisNum]; prs {
			// thisNum was in already
			if history[thisNum][1] != -1 {
				// if it'd been there more than once, bump it
				history[thisNum][0] = history[thisNum][1]
			}
			history[thisNum][1] = loop
		} else {
			// first time thisNum was seen
			history[thisNum] = []int{loop, -1}
		}
		lastNum = thisNum
		loop++
	}
	return lastNum, nil
}
