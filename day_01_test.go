package adventofcode2020

import (
	"testing"
)

func TestDayOne(t *testing.T) {

	testData, err := readFileAsInts("day_01_input.txt")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		input    []int
		num1     int
		num2     int
		expected int64
	}{
		{
			[]int{1721, 979, 366, 299, 675},
			1721,
			299,
			int64(514579),
		},
		{
			testData,
			-1,
			-1,
			-1,
		},
	}
	for _, c := range cases {
		num1, num2, result, err := dayOne(c.input)
		if err != nil {
			t.Fatalf("Error: %s\n", err)
		}
		if c.expected == -1 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("Result: %d + %d = 2020, %d * %d = %d\n",
				num1, num2, num1, num2, result)
			continue
		}
		if num1 != c.num1 || num2 != c.num2 {
			t.Fatalf("Expected %d/%d\nActual%d/%d\n",
				c.num1, c.num2, num1, num2)
		}
		if result != c.expected {
			t.Fatalf("Expected Total %d, Actual %d\n",
				c.expected, result)
		}
	}
}

func TestDayOnePart2(t *testing.T) {

	testData, err := readFileAsInts("day_01_input.txt")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		input    []int
		num1     int
		num2     int
		num3     int
		expected int64
	}{
		{
			[]int{1721, 979, 366, 299, 675},
			979,
			366,
			675,
			int64(241861950),
		},
		{
			testData,
			-1,
			-1,
			-1,
			-1,
		},
	}
	for _, c := range cases {
		num1, num2, num3, result, err := dayOnePart2(c.input)
		if err != nil {
			t.Logf("Error: %s\n", err)
		}
		if c.expected == -1 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("Result: %d + %d = 2020, %d * %d = %d\n",
				num1, num2, num1, num2, result)
			continue
		}
		if num1 != c.num1 || num2 != c.num2 || num3 != c.num3 {
			t.Fatalf("Expected %d/%d/%d\nActual%d/%d/%d\n",
				c.num1, c.num2, c.num3, num1, num2, num3)
		}
		if result != c.expected {
			t.Fatalf("Expected Total %d, Actual %d\n",
				c.expected, result)
		}
	}
}
