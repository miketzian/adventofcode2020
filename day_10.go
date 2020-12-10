package adventofcode2020

import (
	"fmt"
	"sort"
)

// Day10 ...
type Day10 struct {
	data    []int         //  sorted
	cache   map[int][]int // cache index -> subsequent indexes
	results map[int]int64 // results cache - only calculate a (given) index once
}

func (day10 Day10) arrangements() int64 {

	day10.cache = make(map[int][]int)
	day10.results = make(map[int]int64)

	for ix := range day10.data {
		// last one
		if (ix + 1) == len(day10.data) {
			// magic
			day10.cache[ix] = []int{-999}
		} else {
			i2 := ix + 1
			day10.cache[ix] = make([]int, 0)
			for i2 < len(day10.data) && day10.data[i2] <= (day10.data[ix]+3) {
				day10.cache[ix] = append(day10.cache[ix], i2)
				i2++
			}
		}
	}
	// need to add the last one also
	return day10.calculate(0) // , "(0)")
}

func (day10 Day10) calculate(ix int) int64 { // , log string) int64 {
	count := int64(0)
	for _, nextIx := range day10.cache[ix] {
		if nextIx == -999 {
			count++
		} else {
			// count += day10.calculate(nextIx) // , fmt.Sprintf("%s, %d", log, day10.data[ix]))

			if r, prs := day10.results[nextIx]; !prs {
				rc := day10.calculate(nextIx)
				day10.results[nextIx] = rc
				count += rc
			} else {
				count += r
			}

		}
	}
	return count
}

func day10(input []int) (int, int, int64, error) {

	diff1 := 0
	diff2 := 0
	// device counts as one, not listed
	diff3 := 1

	max := 0
	last := 0
	sort.Ints(input)

	for _, v := range input {
		diff := v - last
		if diff > 3 {
			return 0, 0, 0, fmt.Errorf("we have a >3 jump")
		}
		switch diff {
		case 0:
			continue
		case 1:
			diff1++
			last = v
		case 2:
			diff2++
			last = v
		case 3:
			diff3++
			last = v
		}
		if v > max {
			max = v
		}
	}

	// insert a zero start point
	darr := make([]int, 1)
	darr[0] = 0
	for _, v := range input {
		darr = append(darr, v)
	}
	arrangements := Day10{data: darr}.arrangements()

	fmt.Printf("d1=%d, d2=%d, d3=%d\n", diff1, diff2, diff3)

	return max + 3, diff1 * diff3, arrangements, nil
}
