package adventofcode2020

import (
	"strings"
	"testing"
)

func TestDayTwenty(t *testing.T) {

	inputData, err := readFileAsStrings("day_20_input.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Length %d", len(inputData))

	inputTestData, err := readFileAsStrings("day_20_input_test.txt")
	if err != nil {
		t.Error(err)
	}
	t.Logf("[INPUT] Test Data Length %d", len(inputTestData))

	cases := []struct {
		input     []string
		roughness int
		expected  int64
	}{
		{
			inputTestData,
			273,
			20899048083289,
		},
		{
			inputData,
			0,
			0,
		},
	}
	for _, c := range cases {
		result, roughness, err := day20(c.input)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == 0 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("[CALC] Result: %d Roughness %d\n",
				result, roughness)
			continue
		}
		if result != c.expected {
			t.Errorf("[ERROR] Result %d != Expected %d",
				result, c.expected)
		} else if roughness != c.roughness {
			t.Errorf("[ERROR] Roughness %d != Expected %d",
				roughness, c.roughness)
		} else {
			t.Logf("[OK] Result: %d, Roughness %d\n",
				result, roughness)
		}
	}
}

func TestDayTwentyRotateLeft(t *testing.T) {

	tile := &day20tile{1, []string{
		"123",
		"456",
		"789",
	}, []string{"123", "369", "789", "147"}, nil}

	tileStr := strings.Join(tile.data, "\n")

	tile.rotateLeft()
	tileStr2 := strings.Join(tile.data, "\n")

	if tileStr == tileStr2 {
		t.Fatal("rotateLeft didn't rotate data")
	}

	/*
		369
		258
		147
	*/

	tileExp := strings.Join([]string{"369", "258", "147"}, "\n")

	if tileExp != tileStr2 {
		t.Fatalf("rotateLeft didn't work\n%s\n", tileStr2)
	}
	if tile.edges[0] != "369" {
		t.Fatalf("rotateLeft didn't work (top): top=%s, row=%s", tile.edges[0], tile.data[0])
	}
	if tile.edges[1] != "987" {
		t.Fatalf("rotateLeft didn't work (right)")
	}
	if tile.edges[2] != "147" {
		t.Fatalf("rotateLeft didn't work (bottom)")
	}
	if tile.edges[3] != "321" {
		t.Fatalf("rotateLeft didn't work (left)")
	}
}

func TestDayTwentyRotateRight(t *testing.T) {

	tile := &day20tile{1, []string{
		"123",
		"456",
		"789",
	}, []string{"123", "369", "789", "147"}, nil}

	tileStr := strings.Join(tile.data, "\n")

	tile.rotateRight()
	tileStr2 := strings.Join(tile.data, "\n")

	if tileStr == tileStr2 {
		t.Fatal("rotateRight didn't rotate data")
	}

	/*
		741
		852
		963
	*/

	tileExp := strings.Join([]string{"741", "852", "963"}, "\n")

	if tileExp != tileStr2 {
		t.Fatalf("rotateRight didn't work\n%s\n", tileStr2)
	}
	if tile.edges[0] != "741" {
		t.Fatalf("rotateRight didn't work: top=%s, row=%s", tile.edges[0], tile.data[0])
	}
	if tile.edges[1] != "123" {
		t.Fatalf("rotateRight didn't work")
	}
	if tile.edges[2] != "963" {
		t.Fatalf("rotateRight didn't work")
	}
	if tile.edges[3] != "789" {
		t.Fatalf("rotateRight didn't work")
	}
}
