package adventofcode2020

import "testing"

func TestDay17Cycles(t *testing.T) {

	inputData, err := readFileAsStrings("day_17_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Length %d", len(inputData))

	cases := []struct {
		input    []string
		cycles   int
		expected int
	}{
		{
			[]string{".#.", "..#", "###"},
			1,
			11,
		},
		{
			[]string{".#.", "..#", "###"},
			2,
			21,
		},
		{
			[]string{".#.", "..#", "###"},
			3,
			38,
		},
		{
			[]string{".#.", "..#", "###"},
			6,
			112,
		},
		{
			inputData,
			6,
			0,
		},
	}
	for _, c := range cases {
		result, err := day17Cycle(c.input, c.cycles)

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

func TestDay17Cycles2(t *testing.T) {

	inputData, err := readFileAsStrings("day_17_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Length %d", len(inputData))

	cases := []struct {
		input    []string
		cycles   int
		expected int
	}{
		{
			[]string{".#.", "..#", "###"},
			1,
			29,
		},
		{
			[]string{".#.", "..#", "###"},
			2,
			60,
		},
		{
			[]string{".#.", "..#", "###"},
			6,
			848,
		},
		{
			inputData,
			6,
			0,
		},
	}
	for _, c := range cases {
		result, err := day17Cycle2(c.input, c.cycles)

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
