package adventofcode2020

import (
	"testing"
)

func TestDayTwentyThree(t *testing.T) {

	cases := []struct {
		input    string
		moves    int
		expected string
	}{
		{
			"389125467",
			10,
			"92658374",
		},
		{
			"389125467",
			100,
			"67384529",
		},
		{
			"186524973",
			100,
			"",
		},
	}
	for _, c := range cases {
		result, err := day23PartOne(c.input, c.moves)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == "" {
			// we don't know the answer, we're using
			// the computation
			t.Logf("[CALC] Result: %s\n",
				result)
			continue
		}
		if result != c.expected {
			t.Errorf("[ERROR] Result %s != Expected %s",
				result, c.expected)
		} else {
			t.Logf("[OK] Result: %s\n",
				result)
		}
	}
}

func TestDayTwentyThreePart2(t *testing.T) {

	cases := []struct {
		input    string
		moves    int
		expected int64
	}{
		{
			"389125467",
			10000000,
			149245887792,
		},
		{
			"186524973",
			10000000,
			0,
		},
	}
	for _, c := range cases {
		result, err := day23PartTwo(c.input, c.moves)
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
