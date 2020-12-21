package adventofcode2020

import (
	"testing"
)

func TestDayTwentyOne(t *testing.T) {

	inputData, err := readFileAsStrings("day_21_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Length %d", len(inputData))

	cases := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
				"trh fvjkl sbzzf mxmxvkd (contains dairy)",
				"sqjhc fvjkl (contains soy)",
				"sqjhc mxmxvkd sbzzf (contains fish)",
			},
			5,
		},
		{
			inputData,
			0,
		},
	}
	for _, c := range cases {
		result, err := day21(c.input)
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
