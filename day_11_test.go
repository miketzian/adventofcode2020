package adventofcode2020

import (
	"testing"
)

func TestDayEleven(t *testing.T) {

	inputData, err := readFileAsStrings("day_11_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Input Length %d", len(inputData))

	inputTestData, err := readFileAsStrings("day_11_input_test.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Test Input Length %d", len(inputTestData))

	cases := []struct {
		input    []string
		expected int
	}{
		{
			inputTestData,
			37,
		},
		{
			inputData,
			-1,
		},
	}
	for _, c := range cases {
		result, err := day11(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == -1 {
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

func TestDayElevenPart2(t *testing.T) {

	inputData, err := readFileAsStrings("day_11_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Input Length %d", len(inputData))

	inputTestData, err := readFileAsStrings("day_11_input_test.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Test Input Length %d", len(inputTestData))

	cases := []struct {
		input    []string
		expected int
	}{
		{
			inputTestData,
			26,
		},
		{
			inputData,
			-1,
		},
	}
	for _, c := range cases {
		result, err := day11Part2(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == -1 {
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
