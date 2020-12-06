package adventofcode2020

import (
	"sort"
	"testing"
)

func TestDayFive(t *testing.T) {

	inputData, err := readFileAsStrings("day_05_input.txt")
	if err != nil {
		t.Error(err)
	}

	max := -1
	seatIds := make([]int, len(inputData))
	for ix, id := range inputData {
		_, _, seat, err := day5(id)
		if err != nil {
			t.Error(err)
		}
		seatIds[ix] = seat
		if seat > max {
			max = seat
		}
	}
	t.Logf("Max Seat is %d", max)
	sort.Ints(seatIds)

	for ix, v := range seatIds {
		if v < 8 {
			continue
		}
		if seatIds[ix+1] >= v+2 {
			t.Logf("Free Seat %d, -1=%d, +1=%d",
				v+1, v, seatIds[ix+1])
			return
		}
	}
	t.Errorf("Not Found")

}

func TestDayFiveDetermineRow(t *testing.T) {
	expected := 44
	input := "FBFBBFFRLR"
	result, err := day5DetermineRow(input)
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("Incorrect Result [%d] expected [%d] for input [%s]",
			result, expected, input)
	}

}
func TestDayFiveDetermineSeat(t *testing.T) {
	expected := 5
	input := "FBFBBFFRLR"
	result, err := day5DetermineSeat(input)
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("Incorrect Result [%d] expected [%d] for input [%s]",
			result, expected, input)
	}
}

func TestDayFiveSeats(t *testing.T) {
	cases := []struct {
		input string
		row   int
		col   int
		seat  int
	}{
		{"FBFBBFFRLR", 44, 5, 357},
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}
	for _, c := range cases {
		row, col, seat, err := day5(c.input)
		if err != nil {
			t.Fatal(err)
		}

		if c.row != row {
			t.Errorf("Result (%s) row %d != Expected %d",
				c.input, row, c.row)
		}
		if c.col != col {
			t.Errorf("Result (%s) col %d != Expected %d",
				c.input, col, c.col)
		}
		if c.seat != seat {
			t.Errorf("Result (%s) seat %d != Expected %d",
				c.input, seat, c.seat)
		}
	}
}
