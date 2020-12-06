package adventofcode2020

import (
	"testing"
)

func TestDaySix(t *testing.T) {

	inputData, err := readFileAsStrings("day_06_input.txt")
	if err != nil {
		t.Error(err)
	}

	inputTestData, err := readFileAsStrings("day_06_input_test.txt")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		input    []string
		expected int
	}{
		{
			[]string{"abcx", "abcy", "abcz"},
			6,
		},
		{
			inputTestData,
			11,
		},
		{
			inputData,
			-1,
		},
	}
	for _, c := range cases {
		result, err := daySix(c.input)
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

func TestDaySixEveryone(t *testing.T) {

	inputData, err := readFileAsStrings("day_06_input.txt")
	if err != nil {
		t.Error(err)
	}

	inputTestData, err := readFileAsStrings("day_06_input_test.txt")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		input    []string
		expected int
	}{
		{
			[]string{"abcx", "abcy", "abcz"},
			3,
		},
		{
			inputTestData,
			6,
		},
		{
			inputData,
			-1,
		},
	}
	for _, c := range cases {
		result, err := daySixEveryone(c.input)
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
