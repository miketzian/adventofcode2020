package adventofcode2020

import "fmt"

// Day11Seat ...
type Day11Seat struct {
	x         int
	y         int
	state     rune
	adjacents []*Day11Seat
}

// IsEmpty ...
func (seat *Day11Seat) IsEmpty() bool {
	return seat.state == 'L'
}

// OccupiedAdj ...
func (seat *Day11Seat) OccupiedAdj() int {
	occAdj := 0
	for _, adj := range seat.adjacents {
		// fmt.Printf("x=%d, y=%d, ax=%d, ay=%d, adj=%v\n", seat.x, seat.y, adj.x, adj.y, adj.state)
		if adj.state == '#' {
			occAdj++
		}
	}
	// fmt.Printf("x=%d, y=%d, oj=%d\n", seat.x, seat.y, occAdj)
	return occAdj
}

// OccupiedDirections ...
func (seat *Day11Seat) OccupiedDirections(rows [][]rune, maxX int, maxY int) int {
	deltas := [][]int{
		[]int{-1, -1},
		[]int{0, -1},
		[]int{1, -1},
		[]int{-1, 0},
		[]int{1, 0},
		[]int{-1, 1},
		[]int{0, 1},
		[]int{1, 1},
	}

	// check each delta, until the location is either out of range
	// or an occupied or free seat is found
	occDir := 0
	for _, delta := range deltas {
		x := seat.x + delta[0]
		y := seat.y + delta[1]

		if x < 0 || y < 0 || x > maxX || y > maxY {
			continue
		}
		l := rows[y][x]

		for l == '.' {
			// need to apply the delta again
			x += delta[0]
			y += delta[1]

			if x < 0 || y < 0 || x > maxX || y > maxY {
				break
			}
			l = rows[y][x]
		}
		if l == '#' {
			occDir++
		}
	}
	return occDir
}

func day11(input []string) (int, error) {

	maxX := len(input[0]) - 1
	maxY := len(input) - 1

	rows := make([][]rune, len(input))
	for ix, v := range input {
		rows[ix] = []rune(v)
	}

	// populate the
	seats := make(map[string]*Day11Seat)
	numSeats := 0

	isSeat := func(x int, y int) bool {
		if x >= 0 && x <= maxX && y >= 0 && y <= maxY {
			// we're in the right range
			return []rune(input[y])[x] != '.'
		}
		return false
	}

	var getOrMakeSeat func(x int, y int) *Day11Seat

	getOrMakeSeat = func(x int, y int) *Day11Seat {

		code := fmt.Sprintf("%d^^%d", x, y)
		if v, prs := seats[code]; prs {
			// already made, no need to make another
			return v
		}
		// need to make the seat
		// as it may be referenced during the creation of other seats
		seats[code] = &Day11Seat{x, y, rows[y][x], make([]*Day11Seat, 0)}
		numSeats++
		sy := y - 1
		for sy <= y+1 {
			sx := x - 1
			for sy >= 0 && sx <= x+1 {
				if !(sy == y && sx == x) && isSeat(sx, sy) {
					seats[code].adjacents = append(seats[code].adjacents, getOrMakeSeat(sx, sy))
				}
				sx++
			}
			sy++
		}
		return seats[code]
	}

	fmt.Printf("Making Seats\n")

	for y, row := range input {
		for x := range row {
			// skip aisles
			if isSeat(x, y) {
				_ = getOrMakeSeat(x, y)
			}
		}
	}
	fmt.Printf("Created %d seats\n", numSeats)

	loops := 0
	for loops < 1000 {
		changes := make([]*Day11Seat, 0)

		for _, seat := range seats {
			if seat.IsEmpty() && seat.OccupiedAdj() == 0 {
				changes = append(changes, seat)
			}
			if !seat.IsEmpty() && seat.OccupiedAdj() >= 4 {
				changes = append(changes, seat)
			}
		}
		loops++
		if len(changes) == 0 {
			fmt.Printf("Stabilized after %d loops\n", loops)
			break
		} else {
			fmt.Printf("Loop #%d, changes=%d\n", loops, len(changes))
		}
		for _, seat := range changes {
			if seat.IsEmpty() {
				seat.state = '#'
			} else {
				seat.state = 'L'
			}
			rows[seat.y][seat.x] = seat.state
		}
		// to print out seats
		// for _, l := range rows {
		// 	fmt.Println(string(l))
		// }
	}
	occupied := 0
	for _, seat := range seats {
		if !seat.IsEmpty() {
			occupied++
		}
	}
	return occupied, nil
}

func day11Part2(input []string) (int, error) {

	maxX := len(input[0]) - 1
	maxY := len(input) - 1

	rows := make([][]rune, len(input))
	for ix, v := range input {
		rows[ix] = []rune(v)
	}

	// populate the
	seats := make(map[string]*Day11Seat)
	numSeats := 0

	isSeat := func(x int, y int) bool {
		if x >= 0 && x <= maxX && y >= 0 && y <= maxY {
			// we're in the right range
			return []rune(input[y])[x] != '.'
		}
		return false
	}

	var getOrMakeSeat func(x int, y int) *Day11Seat

	getOrMakeSeat = func(x int, y int) *Day11Seat {

		code := fmt.Sprintf("%d^^%d", x, y)
		if v, prs := seats[code]; prs {
			// already made, no need to make another
			return v
		}
		// need to make the seat
		// as it may be referenced during the creation of other seats
		seats[code] = &Day11Seat{x, y, rows[y][x], nil}
		numSeats++
		return seats[code]
	}

	fmt.Printf("Making Seats\n")

	for y, row := range input {
		for x := range row {
			// skip aisles
			if isSeat(x, y) {
				_ = getOrMakeSeat(x, y)
			}
		}
	}
	fmt.Printf("Created %d seats\n", numSeats)

	loops := 0
	for loops < 1000 {
		changes := make([]*Day11Seat, 0)

		for _, seat := range seats {
			occ := seat.OccupiedDirections(rows, maxX, maxY)

			if seat.IsEmpty() && occ == 0 {
				changes = append(changes, seat)
			}
			if !seat.IsEmpty() && occ >= 5 {
				changes = append(changes, seat)
			}
		}
		loops++
		if len(changes) == 0 {
			fmt.Printf("Stabilized after %d loops\n", loops)
			break
		} else {
			fmt.Printf("Loop #%d, changes=%d\n", loops, len(changes))
		}
		for _, seat := range changes {
			if seat.IsEmpty() {
				seat.state = '#'
			} else {
				seat.state = 'L'
			}
			rows[seat.y][seat.x] = seat.state
		}

		for _, l := range rows {
			fmt.Println(string(l))
		}
	}
	occupied := 0
	for _, seat := range seats {
		if !seat.IsEmpty() {
			occupied++
		}
	}
	return occupied, nil
}
