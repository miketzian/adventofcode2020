package adventofcode2020

import (
	"fmt"
)

func day9IsSumOf(input []int64, sumOf int64) bool {
	// very similar to day 1
	for ix, num1 := range input {
		if num1 < sumOf {
			for ix2, num2 := range input {
				if ix2 != ix && num1+num2 == sumOf {
					return true
				}
			}
		}
	}
	return false
}

// find a range of contiguous numbers that sum
// to the sumOf value
// then return min + max
func day9FindRange(input []int64, sumOf int64) (int64, error) {
	last := len(input) - 1

	for ix, v := range input {
		sum := v

		// i sup
		if ix == last {
			if v != sumOf {
				return 0, fmt.Errorf("can't find the number")
			}
			return sum + v, nil
		}

		for ix2, v2 := range input[ix+1:] {
			sum += v2
			if sum == sumOf {
				fmt.Printf("Found index match at %d and %d, %d and %d\n",
					ix, ix2, v, v2)

				min := v
				max := v
				// since ix2 is the index of a slice, we need to account
				// for where that slice started in the max index
				for _, v3 := range input[ix+1 : ix+1+ix2] {
					if v3 < min {
						min = v3
					}
					if v3 > max {
						max = v3
					}
				}
				fmt.Printf("Found min=%d and max=%d\n",
					min, max)
				return min + max, nil
			} else if sum > sumOf {
				break
			}
		}
	}
	return 0, fmt.Errorf("can't find the number")
}

func day9(input []int64, preamble int) (int64, int64, error) {

	for ix, v := range input {
		if ix < preamble {
			// we don't check these
			continue
		}
		min := ix - preamble
		if min < 0 {
			min = 0
		}
		if !day9IsSumOf(input[min:ix], v) {

			weakness, err := day9FindRange(input, v)
			if err != nil {
				return 0, 0, err
			}
			return v, weakness, nil
		}
	}
	return 0, 0, fmt.Errorf("Number not found")
}
