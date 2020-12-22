package adventofcode2020

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type day20tile struct {
	id        int
	data      []string
	edges     []string
	edgeTiles []*day20EdgeMatch
}

func (thisTile *day20tile) top() string {
	return thisTile.edges[0]
}

func (thisTile *day20tile) right() string {
	return thisTile.edges[1]
}

func (thisTile *day20tile) bottom() string {
	return thisTile.edges[2]
}

func (thisTile *day20tile) left() string {
	return thisTile.edges[3]
}

func (thisTile *day20tile) matchRight(thatTile *day20tile) {
	// thisTile doesn't move
	// thatTile moves so that thisTile is on it's left

	side, reverse := thatTile.matchesWith(thisTile.right())
	switch side {
	case 0:
		thatTile.rotateLeft()
	case 1:
		thatTile.flipLeftRight()
	case 2:
		thatTile.rotateRight()
	}
	side, reverse = thatTile.matchesWith(thisTile.right())
	if side != 3 {
		panic("this should be ok by now")
	}
	if reverse {
		thatTile.flipUpDown()
	}
}

func (thisTile *day20tile) matchBelow(thatTile *day20tile) {
	// thisTile doesn't move
	// thatTile moves so that thisTile is on it's left

	side, reverse := thatTile.matchesWith(thisTile.bottom())
	switch side {
	case 1:
		thatTile.rotateLeft()
	case 2:
		thatTile.flipUpDown()
	case 3:
		thatTile.rotateRight()
	}
	side, reverse = thatTile.matchesWith(thisTile.bottom())
	if side != 0 {
		panic("this should be ok by now")
	}
	if reverse {
		thatTile.flipLeftRight()
	}
}

func (thisTile *day20tile) findMonsters() (int, error) {

	// find the orientation in which there are sea monsters.

	/*
		                  #
		#    ##    ##    ###
		 #  #  #  #  #  #
	*/

	line2, err := regexp.Compile("#....##....##....###")
	if err != nil {
		return 0, err
	}
	line3, err := regexp.Compile("#..#..#..#..#..#")
	if err != nil {
		return 0, err
	}

	// if a line matches the regex
	loops := 0

	// outerLoop:
	for loops < 10 {
		monsters := 0
		for ix, line := range thisTile.data {
			if ix == 0 || ix == len(thisTile.data)-1 {
				// we're matching on the 2nd line of the seamonster
				continue
			}

			matches := line2.FindAllStringSubmatchIndex(line, -1)
			if len(matches) > 0 {

				for _, m := range matches {

					// if this is the right spot, line 3 should be
					// one line down, and one character over
					if thisTile.data[ix-1][m[1]-2:m[1]-1] != "#" {
						// the line before doesn't match, so keep iterating
						continue
					}

					line3m := line3.FindStringSubmatchIndex(thisTile.data[ix+1][m[0]+1:])
					// should match at zero as we're basign this
					// on the line above.
					if len(line3m) > 0 && line3m[0] == 0 {
						monsters++
						fmt.Printf("Found Seamonster - loop %d\n", loops)
						// fmt.Println(thisTile.data[ix-1][m[0]:m[1]])
						// fmt.Println(thisTile.data[ix][m[0]:m[1]])
						// fmt.Println(thisTile.data[ix+1][m[0]:m[1]])
					}
				}
			}
		}
		if monsters > 0 {
			// each monster is 15 #s
			roughness := -15 * monsters
			for _, line := range thisTile.data {
				for _, c := range line {
					if c == '#' {
						roughness++
					}
				}
			}
			return roughness, nil
		}
		switch loops {
		case 0:
			thisTile.flipLeftRight()
		case 1:
			thisTile.flipUpDown()
		case 2:
			thisTile.flipLeftRight()
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 5:
			thisTile.rotateLeft()
		case 6:
			thisTile.flipLeftRight()
		case 7:
			thisTile.rotateLeft()
		case 8:
			thisTile.rotateLeft()
		case 9:
			thisTile.rotateLeft()
		}
		loops++
	}
	return 0, fmt.Errorf("did not find the seamonsters")
}

func (thisTile *day20tile) edgesMatched() (bool, bool, bool, bool) {
	// as-is, which edges match

	var top, right, bottom, left bool

	for _, edgeTile := range thisTile.edgeTiles {
		// note, comparisons
		if edgeTile.thisEdge == thisTile.edges[0] || edgeTile.thisEdge == reverseString(thisTile.edges[0]) {
			top = true
		}
		if edgeTile.thisEdge == thisTile.edges[1] || edgeTile.thisEdge == reverseString(thisTile.edges[1]) {
			right = true
		}
		if edgeTile.thisEdge == thisTile.edges[2] || edgeTile.thisEdge == reverseString(thisTile.edges[2]) {
			bottom = true
		}
		if edgeTile.thisEdge == thisTile.edges[3] || edgeTile.thisEdge == reverseString(thisTile.edges[3]) {
			left = true
		}
	}
	return top, right, bottom, left
}

func (thisTile *day20tile) matchesWith(edge string) (int, bool) {

	for ix, e := range thisTile.edges {
		if e == edge {
			return ix, false
		}
		if e == reverseString(edge) {
			return ix, true
		}
	}
	panic("don't call this unless there's a match")
}

func (thisTile *day20tile) rightEdge() *day20EdgeMatch {

	for _, edgeTile := range thisTile.edgeTiles {
		if edgeTile.thisEdge == thisTile.edges[1] || edgeTile.thisEdge == reverseString(thisTile.edges[1]) {
			return edgeTile
		}
	}
	panic("should not be possible")
}

func (thisTile *day20tile) bottomEdge() *day20EdgeMatch {
	for _, edgeTile := range thisTile.edgeTiles {
		if edgeTile.thisEdge == thisTile.edges[2] || edgeTile.thisEdge == reverseString(thisTile.edges[2]) {
			return edgeTile
		}
	}
	panic("should not be possible")
}

func (thisTile *day20tile) rotateLeft() {

	newData := make([]string, len(thisTile.data))
	// re-create the data array
	maxC := len(thisTile.data) - 1

	for rIx, row := range thisTile.data {
		for cIx := range row {
			newData[rIx] += string(thisTile.data[cIx][maxC-rIx])
		}
	}

	// data[2][0] = oldData[0][0] = 1
	// data[1][0] = oldData[0][1] = 2
	// data[0][0] = oldData[0][2] = 3
	// data[2][1] = oldData[1][0] = 4
	// data[1][1] = oldData[1][1] = 5
	// data[0][1] = oldData[1][2] = 6
	// data[2][2] = oldData[2][0] = 7
	// data[1][2] = oldData[2][1] = 8
	// data[0][2] = oldData[2][2] = 9

	// data[0][0] = oldData[0][2] = 3
	// data[0][1] = oldData[1][2] = 6
	// data[0][2] = oldData[2][2] = 9
	// data[1][0] = oldData[0][1] = 2
	// data[1][1] = oldData[1][1] = 5
	// data[1][2] = oldData[2][1] = 8
	// data[2][0] = oldData[0][0] = 1
	// data[2][1] = oldData[1][0] = 4
	// data[2][2] = oldData[2][0] = 7

	thisTile.data = newData

	// top->left, right->top, bottom->right, left->bottom
	if thisTile.edges != nil {
		thisTile.edges[0], thisTile.edges[1], thisTile.edges[2], thisTile.edges[3] =
			thisTile.edges[1], reverseString(thisTile.edges[2]), thisTile.edges[3], reverseString(thisTile.edges[0])
	} else {
		thisTile.makeEdges()
	}
}

func (thisTile *day20tile) rotateRight() {

	newData := make([]string, len(thisTile.data))
	for _, row := range thisTile.data {
		// row[0] should be the last column
		for colIx, col := range row {
			newData[colIx] = string(col) + newData[colIx]
		}
	}
	thisTile.data = newData

	// top->right, right->bottom, bottom->left, left->top
	if thisTile.edges != nil {
		thisTile.edges[0], thisTile.edges[1], thisTile.edges[2], thisTile.edges[3] =
			reverseString(thisTile.edges[3]), thisTile.edges[0], reverseString(thisTile.edges[1]), thisTile.edges[2]
	} else {
		thisTile.makeEdges()
	}
}

func (thisTile *day20tile) flipUpDown() {
	// swap the data
	reverseStringArray(thisTile.data)

	// swap the edges
	if thisTile.edges != nil {
		// left/right
		thisTile.edges[1], thisTile.edges[3] = reverseString(thisTile.edges[1]), reverseString(thisTile.edges[3])
		// top/bottom
		thisTile.edges[0], thisTile.edges[2] = thisTile.edges[2], thisTile.edges[0]
	} else {
		thisTile.makeEdges()
	}
}

func (thisTile *day20tile) flipLeftRight() {
	// swap the data
	for ix, line := range thisTile.data {
		thisTile.data[ix] = reverseString(line)
	}
	// swap the edges
	if thisTile.edges != nil {
		// left/right
		thisTile.edges[1], thisTile.edges[3] = thisTile.edges[3], thisTile.edges[1]
		// top/bottom
		thisTile.edges[0], thisTile.edges[2] = reverseString(thisTile.edges[0]), reverseString(thisTile.edges[2])
	} else {
		thisTile.makeEdges()
	}
}

type day20EdgeMatch struct {
	tile     *day20tile
	thisEdge string
	thatEdge string
}

func (thisTile *day20tile) makeEdges() {
	if thisTile.edges == nil {
		thisTile.edges = make([]string, 4)
	}
	maxX := len(thisTile.data[0]) - 1
	maxY := len(thisTile.data) - 1

	thisTile.edges[0] = thisTile.data[0]
	thisTile.edges[1] = ""
	thisTile.edges[2] = thisTile.data[maxY]
	thisTile.edges[3] = ""
	for _, row := range thisTile.data {
		// left
		thisTile.edges[3] += row[0:1]
		// right
		thisTile.edges[1] += row[maxX:]
	}
}

func (thisTile *day20tile) edgeMatches(otherTiles []*day20tile) int {
	edgeMatches := 0
nextEdge:
	for _, thisEdge := range thisTile.edges {
		for _, thatTile := range otherTiles {
			if thisTile.id == thatTile.id {
				continue
			}
			for _, thatEdge := range thatTile.edges {
				if thisEdge == thatEdge || reverseString(thisEdge) == thatEdge {
					edgeMatches++
					// maybe this should do the reverse as well
					// as this will get calculated in the reverse direction also.
					// thisTile.edgeTiles[ix] = thatTile
					thisTile.edgeTiles = append(thisTile.edgeTiles, &day20EdgeMatch{thatTile, thisEdge, thatEdge})
					// check next edge
					// note, works if we continue or not
					// so suggests there's only one spot for each tile.
					continue nextEdge
				}
			}
		}
	}
	return edgeMatches
}

func day20Load(input []string) ([]*day20tile, error) {

	// data := make(map[int]*day20tile)
	data := make([]*day20tile, 0)

	var thisTile *day20tile

	for _, line := range input {
		if strings.HasPrefix(line, "Tile ") {
			// new tile
			tileID, err := strconv.Atoi(line[5 : len(line)-1])
			if err != nil {
				return nil, err
			}
			thisTile = &day20tile{tileID, nil, nil, nil}
			// data[tileID] = thisTile
			data = append(data, thisTile)
		} else if line == "" {
			// end of tile
			thisTile.makeEdges()
			thisTile = nil
		} else {
			// we have data
			thisTile.data = append(thisTile.data, line)
		}
	}
	if thisTile != nil {
		thisTile.makeEdges()
	}

	return data, nil
}

func day20(input []string) (int64, int, error) {

	data, err := day20Load(input)
	if err != nil {
		return 0, 0, err
	}

	mappedTiles := make(map[int][]*day20tile)

	cornerSum := int64(1)
	for _, thisTile := range data {
		// this will calculate the connecting tiles
		edgeMatches := thisTile.edgeMatches(data)
		// fmt.Printf("Tile %d has edges: %d\n", thisTile.id, edgeMatches)
		if edgeMatches == 2 {
			// this is a corner
			cornerSum *= int64(thisTile.id)
		}
		if edgeMatches != len(thisTile.edgeTiles) {
			// double check
			return 0, 0, fmt.Errorf("edgeMatches didn't match")
		}
		mappedTiles[edgeMatches] = append(mappedTiles[edgeMatches], thisTile)
	}

	if len(mappedTiles[2]) != 4 {
		return 0, 0, fmt.Errorf("we don't have 4 corners")
	}

	gridCalc := math.Sqrt(float64(len(data)))

	if float64(int(gridCalc)) != gridCalc {
		// ie, the number wasn't round
		return 0, 0, fmt.Errorf("we don't know how big the grid is")
	}
	gridSize := int(gridCalc)
	fmt.Printf("Grid is %d on a side\n", gridSize)

	grid := make([][]*day20tile, gridSize)

	thisTile := mappedTiles[2][0]
	grid[0] = make([]*day20tile, gridSize)
	grid[0][0] = thisTile
	// start with a corner, any corner

	tileSize := len(thisTile.data)

	top, right, bottom, left := thisTile.edgesMatched()
	if !right {
		thisTile.flipLeftRight()
	}
	if !bottom {
		thisTile.flipUpDown()
	}
	top, right, bottom, left = thisTile.edgesMatched()
	if !(!top && !left && right && bottom) {
		return 0, 0, fmt.Errorf("we have a logic error")
	}

	fmt.Printf("Top Left Is: %d\n", grid[0][0].id)

	rowIx := 0
	for rowIx < gridSize {

		// lastRow := rowIx == gridSize-1
		if rowIx > 0 {
			fmt.Printf("Processing %d/0 .. ", rowIx)
			grid[rowIx] = make([]*day20tile, gridSize)
			// match the one to the bottom of the top left

			tileAbove := grid[rowIx-1][0]
			tileBelow := tileAbove.bottomEdge()
			// rotate tileBelow into position
			tileAbove.matchBelow(tileBelow.tile)
			grid[rowIx][0] = tileBelow.tile

			fmt.Println(grid[rowIx][0].id)
		}
		colIx := 1
		for colIx < gridSize {
			fmt.Printf("Processing %d/%d.. ", rowIx, colIx)
			tileLeft := grid[rowIx][colIx-1]
			tileRight := tileLeft.rightEdge()
			tileLeft.matchRight(tileRight.tile)
			grid[rowIx][colIx] = tileRight.tile
			fmt.Println(grid[rowIx][colIx].id)
			colIx++
		}
		rowIx++
	}

	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%d ", col.id)
		}
		fmt.Println()
	}
	// now we need to turn the grid into one big tile

	lines := make([]string, 0)

	for _, row := range grid {
		lineIx := 1
		for lineIx < (tileSize - 1) {
			line := ""
			for _, col := range row {
				line += col.data[lineIx][1 : tileSize-1]
			}
			lineIx++
			lines = append(lines, line)
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	finalTile := &day20tile{1, lines, nil, nil}

	roughSeas, err := finalTile.findMonsters()

	if err != nil {
		return 0, 0, err
	}

	return cornerSum, roughSeas, nil
}
