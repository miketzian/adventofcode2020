package adventofcode2020

import (
	"testing"
)

func TestDayTen(t *testing.T) {

	inputData, err := readFileAsInts("day_10_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Input Length %d", len(inputData))
	inputTestData, err := readFileAsInts("day_10_input_test.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Test Input Length %d", len(inputTestData))

	cases := []struct {
		input        []int
		max          int
		expected     int
		arrangements int64
	}{
		{
			[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4},
			22,
			35,
			8,
		},
		{
			inputTestData,
			-1,
			220,
			19208,
		},
		{
			inputData,
			-1,
			-1,
			-1,
		},
	}
	for _, c := range cases {
		_, result, arrangements, err := day10(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == -1 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("[CALC] Result: %d, Arrangements: %d\n",
				result, arrangements)
			continue
		}
		if result != c.expected || arrangements != c.arrangements {
			t.Errorf("[ERROR] Result %d != Expected %d || arrangements %d != Expected %d",
				result, c.expected, arrangements, c.arrangements)
		} else {
			t.Logf("[OK] Result: %d, Arrangements: %d\n",
				result, arrangements)
		}
	}
}
