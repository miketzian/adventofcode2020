package adventofcode2020

import (
	"testing"
)

func TestDayThirteen(t *testing.T) {

	inputData, err := readFileAsStrings("day_13_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Input Length %d", len(inputData))

	inputTestData, err := readFileAsStrings("day_13_input_test.txt")
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
			295,
		},
		{
			inputData,
			-1,
		},
	}
	for _, c := range cases {
		result, err := day13(c.input)
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

func TestDayThirteenPart2(t *testing.T) {

	inputData, err := readFileAsStrings("day_13_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Input Length %d", len(inputData))

	inputTestData, err := readFileAsStrings("day_13_input_test.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Test Input Length %d", len(inputTestData))

	cases := []struct {
		input    string
		expected uint64
	}{
		{
			"17,x,13,19",
			3417,
		},
		{
			"67,7,59,61",
			754018,
		},
		{
			"67,x,7,59,61",
			779210,
		},
		{
			"67,7,x,59,61",
			1261476,
		},
		{
			"1789,37,47,1889",
			1202161486,
		},
		{
			inputTestData[1],
			1068781,
		},
		{
			inputData[1],
			0,
		},
	}
	for _, c := range cases {
		result, err := day13Part2(c.input)
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
