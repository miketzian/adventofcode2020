package adventofcode2020

import (
	"fmt"
	"strconv"
)

// Day12 ...
type Day12 struct {
	x   int
	y   int
	dir int
	wpx int
	wpy int
}

// Move ...
func (day *Day12) Move(move string) error {

	action := move[0]
	num, err := strconv.Atoi(move[1:])
	if err != nil {
		return err
	}

	if action == 'F' {
		switch day.dir {
		case 0:
			action = 'N'
		case 90:
			action = 'E'
		case 180:
			action = 'S'
		case 270:
			action = 'W'
		default:
			return fmt.Errorf("Unexpected dir %d", day.dir)
		}
	}

	switch action {
	case 'N':
		day.y += num
	case 'S':
		day.y -= num
	case 'E':
		day.x += num
	case 'W':
		day.x -= num
	case 'L':
		day.dir -= num
		if day.dir < 0 {
			day.dir += 360
		}
	case 'R':
		day.dir += num
		if day.dir >= 360 {
			day.dir -= 360
		}
	case 'F':
		fallthrough
	default:
		return fmt.Errorf("Unexpected action %s", string(action))
	}
	return nil
}

// ManhattanDistance ...
func (day *Day12) ManhattanDistance() int {
	x := day.x
	y := day.y

	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}

func day12(input []string) (int, error) {

	day := &Day12{0, 0, 90, 0, 0}

	for _, instruction := range input {
		err := day.Move(instruction)
		if err != nil {
			return 0, err
		}
	}
	return day.ManhattanDistance(), nil
}

// Operate ...
func (day *Day12) Operate(move string) error {

	action := move[0]
	num, err := strconv.Atoi(move[1:])
	if err != nil {
		return err
	}

	if action == 'F' {
		day.x += (day.wpx * num)
		day.y += (day.wpy * num)
		return nil
	}

	switch action {
	case 'N':
		day.wpy += num
	case 'S':
		day.wpy -= num
	case 'E':
		day.wpx += num
	case 'W':
		day.wpx -= num
	case 'L':
		wpx := day.wpx
		wpy := day.wpy

		// N3 -> 10E, 4N
		// R90 -> 4E, 10S

		switch num {
		case 90:
			// turn left 90
			day.wpx = (wpy * -1)
			day.wpy = wpx
		case 180:
			day.wpx *= -1
			day.wpy *= -1
		case 270:
			day.wpx = wpy
			day.wpy = (wpx * -1)
		default:
			return fmt.Errorf("Unexpected turn %d", num)
		}
	case 'R':
		wpx := day.wpx
		wpy := day.wpy
		switch num {
		case 90:
			day.wpx = wpy
			day.wpy = (wpx * -1)
		case 180:
			day.wpx *= -1
			day.wpy *= -1
		case 270:
			day.wpx = (wpy * -1)
			day.wpy = wpx
		default:
			return fmt.Errorf("Unexpected turn %d", num)
		}
	case 'F':
		fallthrough
	default:
		return fmt.Errorf("Unexpected action %s", string(action))
	}
	return nil
}

func day12Part2(input []string) (int, error) {

	day := &Day12{0, 0, 90, 10, 1}

	for _, instruction := range input {
		err := day.Operate(instruction)
		if err != nil {
			return 0, err
		}
		// fmt.Printf("Op: %s, Ship: e=%d, n=%d, wpe=%d, wpn=%d\n",
		// 	instruction, day.x, day.y, day.wpx, day.wpy)
	}
	return day.ManhattanDistance(), nil
}
