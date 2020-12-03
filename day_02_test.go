package adventofcode2020

import (
	"testing"
)

func TestDayTwo(t *testing.T) {

	inputData, err := readFileAsStrings("day_02_input.txt")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		input    []string
		expected int
	}{
		{
			[]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"},
			2,
		},
		{
			inputData,
			-1,
		},
	}
	for _, c := range cases {
		result, err := dayTwo(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == -1 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("Result: %d\n", result)
			continue
		}
		if result != c.expected {
			t.Errorf("Result %d != Expected %d",
				result, c.expected)
		}
	}
}

func TestDayTwoPart2(t *testing.T) {

	inputData, err := readFileAsStrings("day_02_input.txt")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		input    []string
		expected int
	}{
		{
			[]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"},
			1,
		},
		{
			inputData,
			-1,
		},
	}
	for _, c := range cases {
		result, err := dayTwoPart2(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == -1 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("Result: %d\n", result)
			continue
		} else if result != c.expected {
			t.Errorf("Result %d != Expected %d",
				result, c.expected)
		}
	}
}
