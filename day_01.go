package adventofcode2020

import (
	"fmt"
)

func dayOne(nums []int) (int, int, int64, error) {

	for ix, num1 := range nums {
		if num1 < 2020 {
			for ix2, num2 := range nums {
				if ix2 != ix && num1+num2 == 2020 {
					return num1, num2, int64(num1) * int64(num2), nil
				}
			}
		}
	}
	return 0, 0, 0, fmt.Errorf("Unable to find the pair")
}

func dayOnePart2(nums []int) (int, int, int, int64, error) {

	for ix, num1 := range nums {
		if num1 < 2020 {
			for ix2, num2 := range nums {
				if ix2 != ix && num1+num2 < 2020 {
					for ix3, num3 := range nums {
						if ix3 != ix && ix3 != ix2 && num1+num2+num3 == 2020 {
							return num1, num2, num3, int64(num1) * int64(num2) * int64(num3), nil
						}
					}
				}
			}
		}
	}
	return 0, 0, 0, 0, fmt.Errorf("Unable to find the pair")
}
