package adventofcode2020

import (
	"strings"
	"testing"
)

func TestDayNineteen(t *testing.T) {

	inputData, err := readFileAsStrings("day_19_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Length %d", len(inputData))

	inputTestData, err := readFileAsStrings("day_19_input_test.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Test Data Length %d", len(inputTestData))

	inputTestData2, err := readFileAsStrings("day_19_input_test_part2.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Test Part2 Data Length %d", len(inputTestData2))

	cases := []struct {
		input    []string
		part2    bool
		expected int
	}{
		{
			inputTestData,
			false,
			2,
		},
		{
			inputData,
			false,
			0,
		},
		{
			inputTestData2,
			false,
			3,
		},
		{
			// simpler case for part 2
			// num | num num2
			strings.Split(`0: 3 4 2 3
2: 5 | 5 6
3: "c"
4: "d"
5: "e"
6: "f"

cdefc`, "\n"),
			true,
			1,
		},
		{
			inputTestData2,
			true,
			12,
		},
		{
			inputData,
			true,
			0,
		},
	}
	for _, c := range cases {
		result, err := day19(c.input, c.part2)
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
