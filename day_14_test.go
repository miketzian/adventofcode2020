package adventofcode2020

import (
	"testing"
)

func TestDayFourteen(t *testing.T) {

	inputData, err := readFileAsStrings("day_14_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Length %d", len(inputData))

	inputTestData, err := readFileAsStrings("day_14_input_test.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Test Data Length %d", len(inputTestData))

	cases := []struct {
		input    []string
		expected uint64
	}{
		{
			inputTestData,
			165,
		},
		{
			inputData,
			0,
		},
	}
	for _, c := range cases {
		result, err := day14(c.input)
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

func TestDayFourteenPart2(t *testing.T) {

	inputData, err := readFileAsStrings("day_14_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Length %d", len(inputData))

	inputTestData, err := readFileAsStrings("day_14_input_test_part2.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Test Data Length %d", len(inputTestData))

	cases := []struct {
		input    []string
		expected uint64
	}{
		{
			inputTestData,
			208,
		},
		{
			inputData,
			0,
		},
	}
	for _, c := range cases {
		result, err := day14Part2(c.input)
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
