package adventofcode2020

import "fmt"

type day17change struct {
	x     int
	y     int
	z     int
	state rune
}

func (o *day17change) String() string {
	return fmt.Sprintf("z=%d, y=%d, x=%d", o.z, o.y, o.x)
}

func day17Cycle(input []string, cycles int) (int, error) {

	// active=#
	// inactive=.

	// cycle through the

	// we always start with one plane
	minZ := 0
	maxZ := 0

	// start at zero
	minY := 0
	minX := 0

	// we will calculate these
	maxY := 0
	maxX := 0

	key := func(x int, y int, z int) string {
		return fmt.Sprintf("%d__%d__%d", z, x, y)
	}

	state := make(map[string]rune)

	for y, line := range input {
		for x, r := range line {
			if r == '#' {
				state[key(x, y, 0)] = r
			}
			if x > maxX {
				maxX = x
			}
		}
		if y > maxY {
			maxY = y
		}
	}

	isActive := func(x int, y int, z int) bool {
		if v, prs := state[key(x, y, z)]; prs {
			return v == '#'
		}
		return false
	}

	setActive := func(x int, y int, z int) {
		state[key(x, y, z)] = '#'
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
		if z < minZ {
			minZ = z
		}
		if z > maxZ {
			maxZ = z
		}
	}

	setInactive := func(x int, y int, z int) {
		delete(state, key(x, y, z))
		// state[key(x, y, z)] = '.'
	}

	activeNeighbours := func(x int, y int, z int) int {
		active := 0
		for nz := z - 1; nz <= z+1; nz++ {
			for ny := y - 1; ny <= y+1; ny++ {
				for nx := x - 1; nx <= x+1; nx++ {
					if nx == x && ny == y && nz == z {
						continue
					}
					if isActive(nx, ny, nz) {
						active++
					}
				}
			}
		}
		return active
	}

	loops := 0
	for loops < cycles {
		loops++
		changes := make([]*day17change, 0)

		for z := minZ - 1; z <= maxZ+1; z++ {
			for y := minY - 1; y <= maxY+1; y++ {
				for x := minX - 1; x <= maxX+1; x++ {
					activeNeighbours := activeNeighbours(x, y, z)
					if isActive(x, y, z) {
						// exists, and is active
						if activeNeighbours != 2 && activeNeighbours != 3 {
							changes = append(changes, &day17change{x, y, z, '.'})
						}
					} else if activeNeighbours == 3 {
						changes = append(changes, &day17change{x, y, z, '#'})
					}
				}
			}
		}
		if len(changes) == 0 {
			fmt.Printf("No more changes on loop %d\n", loops)
		}
		for _, c := range changes {
			if c.state == '#' {
				setActive(c.x, c.y, c.z)
			} else {
				// fmt.Printf("Set %s inactive\n", c.String())
				setInactive(c.x, c.y, c.z)
			}
		}

		fmt.Printf("Loop %d\n", loops)
		/*
			loopActive = 0
			for z := minZ; z <= maxZ; z++ {
				fmt.Printf("z=%d\n", z)
				for y := minY; y <= maxY; y++ {
					for x := minX; x <= maxX; x++ {
						if isActive(x, y, z) {
							loopActive++
							fmt.Printf("#")
						} else {
							fmt.Printf(".")
						}
					}
					fmt.Println()
				}
			}
			fmt.Println(state)
		*/
	}

	return len(state), nil
}

type day17change2 struct {
	x     int
	y     int
	z     int
	w     int
	state rune
}

func (o *day17change2) String() string {
	return fmt.Sprintf("w=%d, z=%d, y=%d, x=%d", o.w, o.z, o.y, o.x)
}

func day17Cycle2(input []string, cycles int) (int, error) {

	// active=#
	// inactive=.

	// cycle through the

	// we always start with one plane
	minZ := 0
	maxZ := 0

	minW := 0
	maxW := 0

	// start at zero
	minY := 0
	minX := 0

	// we will calculate these
	maxY := 0
	maxX := 0

	key := func(x int, y int, z int, w int) string {
		return fmt.Sprintf("%d__%d__%d__%d", w, z, x, y)
	}

	state := make(map[string]rune)

	for y, line := range input {
		for x, r := range line {
			if r == '#' {
				state[key(x, y, 0, 0)] = r
			}
			if x > maxX {
				maxX = x
			}
		}
		if y > maxY {
			maxY = y
		}
	}

	isActive := func(x int, y int, z int, w int) bool {
		if v, prs := state[key(x, y, z, w)]; prs {
			return v == '#'
		}
		return false
	}

	setActive := func(x int, y int, z int, w int) {
		state[key(x, y, z, w)] = '#'
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
		if z < minZ {
			minZ = z
		}
		if z > maxZ {
			maxZ = z
		}
		if w < minW {
			minW = w
		}
		if w > maxW {
			maxW = w
		}
	}

	setInactive := func(x int, y int, z int, w int) {
		delete(state, key(x, y, z, w))
	}

	activeNeighbours := func(x int, y int, z int, w int) int {
		active := 0
		for nw := w - 1; nw <= w+1; nw++ {
			for nz := z - 1; nz <= z+1; nz++ {
				for ny := y - 1; ny <= y+1; ny++ {
					for nx := x - 1; nx <= x+1; nx++ {
						if nx == x && ny == y && nz == z && nw == w {
							continue
						}
						if isActive(nx, ny, nz, nw) {
							active++
						}
					}
				}
			}
		}
		return active
	}

	loops := 0
	for loops < cycles {
		loops++
		changes := make([]*day17change2, 0)

		for w := minW - 1; w <= maxW+1; w++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				for y := minY - 1; y <= maxY+1; y++ {
					for x := minX - 1; x <= maxX+1; x++ {
						activeNeighbours := activeNeighbours(x, y, z, w)
						if isActive(x, y, z, w) {
							// exists, and is active
							if activeNeighbours != 2 && activeNeighbours != 3 {
								changes = append(changes, &day17change2{x, y, z, w, '.'})
							}
						} else if activeNeighbours == 3 {
							changes = append(changes, &day17change2{x, y, z, w, '#'})
						}
					}
				}
			}
		}
		if len(changes) == 0 {
			fmt.Printf("No more changes on loop %d\n", loops)
		}
		for _, c := range changes {
			if c.state == '#' {
				setActive(c.x, c.y, c.z, c.w)
			} else {
				// fmt.Printf("Set %s inactive\n", c.String())
				setInactive(c.x, c.y, c.z, c.w)
			}
		}

		fmt.Printf("Loop %d\n", loops)
		/*
			loopActive = 0
			for z := minZ; z <= maxZ; z++ {
				fmt.Printf("z=%d\n", z)
				for y := minY; y <= maxY; y++ {
					for x := minX; x <= maxX; x++ {
						if isActive(x, y, z) {
							loopActive++
							fmt.Printf("#")
						} else {
							fmt.Printf(".")
						}
					}
					fmt.Println()
				}
			}
			fmt.Println(state)
		*/
	}

	return len(state), nil
}
