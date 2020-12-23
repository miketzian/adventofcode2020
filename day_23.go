package adventofcode2020

import (
	"container/ring"
	"fmt"
	"strconv"
)

func day23PartOne(input string, moves int) (string, error) {

	cups := ring.New(len(input))
	cupLookup := make(map[int]*ring.Ring)

	for _, r := range input {
		num, err := strconv.Atoi(string(r))
		if err != nil {
			return "", err
		}
		cups.Value = num
		cupLookup[num] = cups
		cups = cups.Next()
	}

	cups, err := day23(cups, cupLookup, moves)
	if err != nil {
		return "", err
	}

	one := cupLookup[1].Next()
	label := ""
	one.Do(func(value interface{}) {
		if value.(int) != 1 {
			label += fmt.Sprintf("%d", value.(int))
		}
	})
	return label, nil
}

func day23PartTwo(input string, moves int) (int64, error) {
	cups := ring.New(1000000)
	cupLookup := make(map[int]*ring.Ring)
	max := -1
	for _, r := range input {
		num, err := strconv.Atoi(string(r))
		if err != nil {
			return 0, err
		}
		cups.Value = num
		cupLookup[num] = cups
		if num > max {
			max = num
		}
		cups = cups.Next()
	}
	max++
	for max <= 1000000 {
		cups.Value = max
		cupLookup[max] = cups
		cups = cups.Next()
		max++
	}
	cups, err := day23(cups, cupLookup, moves)
	one := cupLookup[1]
	first := one.Next().Value.(int)
	second := one.Next().Next().Value.(int)
	sum := int64(first) * int64(second)
	fmt.Printf("Values: %d * %d = %d\n", first, second, sum)
	return sum, err
}

func day23(cups *ring.Ring, cupLookup map[int]*ring.Ring, moves int) (*ring.Ring, error) {

	currentCup := cups

	min := func(loop *ring.Ring) int {
		min := loop.Value.(int)
		loop.Do(func(value interface{}) {
			if value.(int) < min {
				min = value.(int)
			}
		})
		return min
	}
	max := func(loop *ring.Ring) int {
		max := loop.Value.(int)
		loop.Do(func(value interface{}) {
			if value.(int) > max {
				max = value.(int)
			}
		})
		return max
	}

	search := func(loop *ring.Ring, value int) *ring.Ring {
		if value == loop.Value.(int) {
			return loop
		}
		for p := loop.Next(); p != loop; p = p.Next() {
			if value == p.Value.(int) {
				return p
			}
		}
		return nil
	}

	// printCups := func(loop *ring.Ring) {
	// 	c := 0
	// 	for c < loop.Len() {
	// 		if loop == currentCup {
	// 			fmt.Printf(" (%d)", loop.Value.(int))
	// 		} else {
	// 			fmt.Printf(" %d", loop.Value.(int))
	// 		}
	// 		loop = loop.Next()
	// 		c++
	// 	}
	// }

	move := 1
	for move <= moves {
		// fmt.Printf("-- move %d --\n", move)
		// fmt.Printf("cups:")
		// printCups(cups)
		// fmt.Println()

		nextThree := currentCup.Unlink(3)

		// fmt.Printf("pick up:")
		// printCups(nextThree)
		// fmt.Println()

		destinationCup := currentCup.Value.(int) - 1
		for search(nextThree, destinationCup) != nil {
			destinationCup--
		}
		// if cups.Len() == 1000000 {
		if moves == 10000000 {
			// we know the len and the min value
			// the max might be in the 3, so we need to account for that
			if destinationCup < 1 {
				destinationCup = 1000000
				for search(nextThree, destinationCup) != nil {
					destinationCup--
				}
			}
		} else {
			if destinationCup < min(cups) {
				destinationCup = max(cups)
			}
		}
		cups = cupLookup[destinationCup].Link(nextThree)
		currentCup = currentCup.Next()

		move++
	}
	return cups, nil
}
