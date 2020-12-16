package adventofcode2020

import (
	"testing"
)

func TestDaySixteen(t *testing.T) {

	inputData, err := readFileAsStrings("day_16_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Input Length %d", len(inputData))

	inputTestData, err := readFileAsStrings("day_16_input_test.txt")
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
			71,
		},
		{
			inputData,
			-1,
		},
	}
	for _, c := range cases {
		result, err := day16(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == -1 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("[CALC] Result: %d \n",
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

//
// class = 1,2
// row = 0,1
// seat = 2

func TestDaySixteenFieldMap(t *testing.T) {

	inputTestData, err := readFileAsStrings("day_16_input_test_map.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Test Input Length %d", len(inputTestData))

	cases := []struct {
		input    []string
		expected map[string]int
	}{
		{
			inputTestData,
			map[string]int{
				"class": 12,
				"row":   11,
				"seat":  13,
			},
		},
	}
	for _, c := range cases {
		result, err := day16MapFields(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == nil {
			// we don't know the answer, we're using
			// the computation
			t.Logf("[CALC] Result: %v \n",
				result)
			continue
		}
		for k, v := range c.expected {
			if result[k] != v {
				t.Errorf("[ERROR] Result %d != Expected %d",
					result[k], v)
			}
		}
		t.Logf("[OK] Result: %v\n", result)
	}
}

func TestDaySixteenPart2(t *testing.T) {

	inputData, err := readFileAsStrings("day_16_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Input Length %d", len(inputData))

	cases := []struct {
		input    []string
		expected int
	}{
		{
			inputData,
			-1,
		},
	}
	for _, c := range cases {
		result, err := day16Part2(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == -1 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("[CALC] Result: %d \n",
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
