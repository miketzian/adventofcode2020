package adventofcode2020

import (
	"testing"
)

func TestDayTwentyTwo(t *testing.T) {

	inputData, err := readFileAsStrings("day_22_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Length %d", len(inputData))

	cases := []struct {
		input    []string
		expected int
	}{
		{
			[]string{"Player 1:", "9", "2", "6", "3", "1", "",
				"Player 2:", "5", "8", "4", "7", "10"},
			306,
		},
		{
			inputData,
			0,
		},
	}
	for _, c := range cases {
		result, err := day22(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == 0 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("[CALC] Result: %d\n",
				result)
			continue
		}
		if result != c.expected {
			t.Errorf("[ERROR] Result %d != Expected %d",
				result, c.expected)
		} else {
			t.Logf("[OK] Result: %d\n",
				result)
		}
	}
}

func TestDayTwentyTwoPart2(t *testing.T) {

	inputData, err := readFileAsStrings("day_22_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Length %d", len(inputData))

	cases := []struct {
		input    []string
		expected int
	}{
		{
			[]string{"Player 1:", "9", "2", "6", "3", "1", "",
				"Player 2:", "5", "8", "4", "7", "10"},
			291,
		},
		{
			inputData,
			0,
		},
	}
	for _, c := range cases {
		result, err := day22Part2(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == 0 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("[CALC] Result: %d\n",
				result)
			continue
		}
		if result != c.expected {
			t.Errorf("[ERROR] Result %d != Expected %d",
				result, c.expected)
		} else {
			t.Logf("[OK] Result: %d\n",
				result)
		}
	}
}
