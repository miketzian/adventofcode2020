package adventofcode2020

// DayThreeResult ...
//type DayThreeResult struct {
//	actual int
//}

func dayThree(input []string, right int, down int) (int, error) {

	x, y := 0, 0
	trees := 0

	for y < len(input) {
		if input[y][x] == '#' {
			trees++
		}
		x = (x + right) % len(input[0])
		y = y + down
	}
	return trees, nil
}

// 211

// Right 1, down 1.  67
// Right 3, down 1. (This is the slope you already checked.)
// Right 5, down 1. 77
// Right 7, down 1. 89
// Right 1, down 2.
