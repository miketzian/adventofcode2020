package adventofcode2020

import (
	"testing"
)

func TestDaySeven(t *testing.T) {

	inputData, err := readFileAsStrings("day_07_input.txt")
	if err != nil {
		t.Error(err)
	}

	inputTestData, err := readFileAsStrings("day_07_input_test.txt")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		input         []string
		expected      int
		expectedCount int
	}{
		{
			inputTestData,
			4,
			32,
		},
		{
			inputData,
			-1,
			-1,
		},
	}
	for _, c := range cases {
		result, resultCount, err := daySeven(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == -1 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("Result: %d, Total: %d\n", result, resultCount)
			continue
		}
		if result != c.expected {
			t.Errorf("Result %d != Expected %d",
				result, c.expected)
		}
		if resultCount != c.expectedCount {
			t.Errorf("Result Count %d != Expected Count %d",
				resultCount, c.expectedCount)
		}
	}
}
