package adventofcode2020

import (
	"testing"
)

func TestDayEight(t *testing.T) {

	inputData, err := readFileAsStrings("day_08_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Input Length %d", len(inputData))
	inputTestData, err := readFileAsStrings("day_08_input_test.txt")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		input    []string
		expected int
	}{
		{
			inputTestData,
			5,
		},
		{
			inputData,
			-1,
		},
	}
	for _, c := range cases {
		result, err := dayEight(c.input)
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

func TestDayEightPart2(t *testing.T) {

	inputData, err := readFileAsStrings("day_08_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Input Length %d", len(inputData))
	inputTestData, err := readFileAsStrings("day_08_input_test.txt")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		input    []string
		expected int
	}{
		{
			inputTestData,
			8,
		},
		{
			inputData,
			-1,
		},
	}
	for _, c := range cases {
		result, err := dayEightPart2(c.input)
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
