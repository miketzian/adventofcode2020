package adventofcode2020

import "testing"

func TestDay18(t *testing.T) {

	inputData, err := readFileAsStrings("day_18_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Length %d", len(inputData))

	cases := []struct {
		input    []string
		expected int
	}{
		{[]string{"1 + 2 * 3 + 4 * 5 + 6"}, 71},
		{[]string{"1 + (2 * 3) + (4 * (5 + 6))"}, 51},
		{[]string{"2 * 3 + (4 * 5)"}, 26},
		{[]string{"5 + (8 * 3 + 9 + 3 * 4 * 3)"}, 437},
		{[]string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}, 12240},
		{[]string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}, 13632},
		{inputData, 0},
	}
	for _, c := range cases {
		result, err := day18(c.input, false)

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

func TestDay18Line(t *testing.T) {

	inputData, err := readFileAsStrings("day_18_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Length %d", len(inputData))

	cases := []struct {
		input    string
		part2    bool
		expected int
	}{
		{"1 + 1", false, 2},
		{"1 + 1 * 6", false, 12},
		{"1 + 1 * 6", true, 12},
		{"2 * 6 + 1", false, 13},
		{"2 * 6 + 1", true, 14},
	}
	for _, c := range cases {
		result, err := day18Line(c.input, c.part2)

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

func TestDay18Part2(t *testing.T) {

	inputData, err := readFileAsStrings("day_18_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Length %d", len(inputData))

	cases := []struct {
		input    []string
		expected int
	}{
		{[]string{"1 + 2 * 3 + 4 * 5 + 6"}, 231},
		{[]string{"1 + (2 * 3) + (4 * (5 + 6))"}, 51},
		{[]string{"2 * 3 + (4 * 5)"}, 46},
		{[]string{"5 + (8 * 3 + 9 + 3 * 4 * 3)"}, 1445},
		{[]string{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"}, 669060},
		{[]string{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"}, 23340},
		{inputData, 0},
	}
	for _, c := range cases {
		result, err := day18(c.input, true)

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
