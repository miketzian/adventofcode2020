package adventofcode2020

import (
	"testing"
)

func TestDayNine(t *testing.T) {

	inputData, err := readFileAsInt64s("day_09_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Input Length %d", len(inputData))
	inputTestData, err := readFileAsInt64s("day_09_input_test.txt")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		input    []int64
		preamble int
		expected int64
		weakness int64
	}{
		{
			inputTestData,
			5,
			127,
			62,
		},
		{
			inputData,
			25,
			-1,
			-1,
		},
	}
	for _, c := range cases {
		result, weakness, err := day9(c.input, c.preamble)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == -1 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("Result: %d, Weakness: %d\n",
				result, weakness)
			continue
		}
		if result != c.expected {
			t.Errorf("Result %d != Expected %d",
				result, c.expected)
		}
		if weakness != c.weakness {
			t.Errorf("Weakness %d != Expected %d",
				weakness, c.weakness)
		}
	}
}
