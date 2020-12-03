package adventofcode2020

import (
	"testing"
)

func TestDayThree(t *testing.T) {

	inputData, err := readFileAsStrings("day_03_input.txt")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		input    []string
		right    int
		down     int
		expected int
	}{
		{
			[]string{"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#"},
			3,
			1,
			7,
		},
		{
			inputData,
			1,
			1,
			-1,
		},
		{
			inputData,
			3,
			1,
			-1,
		},
		{
			inputData,
			5,
			1,
			-1,
		},
		{
			inputData,
			7,
			1,
			-1,
		},
		{
			inputData,
			1,
			2,
			-1,
		},
	}
	mult := 1
	for _, c := range cases {
		result, err := dayThree(c.input, c.right, c.down)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == -1 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("Down: %d, Right: %d, Result: %d\n",
				c.down, c.right, result)
			mult *= result
			continue
		}
		if result != c.expected {
			t.Errorf("Result %d != Expected %d",
				result, c.expected)
		}
	}
	t.Logf("Result: %d\n", mult)
}
