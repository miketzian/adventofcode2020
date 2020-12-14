package adventofcode2020

import (
	"fmt"
	"regexp"
	"strconv"
)

func day14(input []string) (uint64, error) {

	maskRe := regexp.MustCompile(`^mask = ([X10]+)$`)
	memRe := regexp.MustCompile(`^mem\[([0-9]+)\] = ([0-9]+)$`)

	setBit := func(n uint64, pos uint) uint64 {
		n |= (1 << pos)
		return n
	}
	clearBit := func(n uint64, pos uint) uint64 {
		mask := ^(1 << pos)
		n &= uint64(mask)
		return n
	}
	reverse := func(s string) string {
		rns := []rune(s) // convert to rune
		for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

			// swap the letters of the string,
			// like first with last and so on.
			rns[i], rns[j] = rns[j], rns[i]
		}
		// return the reversed string.
		return string(rns)
	}

	mem := make(map[int]uint64)
	mask := ""
	for _, line := range input {

		// either the line will be a pattern

		// or the line will be an assignment
		match := maskRe.FindStringSubmatch(line)
		if match != nil {
			// we have a new mask
			mask = reverse(match[1])
			continue
		}
		match = memRe.FindStringSubmatch(line)
		if match != nil {
			// we have a new memory line
			addr, err := strconv.Atoi(match[1])
			if err != nil {
				return 0, err
			}
			value, err := strconv.ParseUint(match[2], 10, 64)
			if err != nil {
				return 0, err
			}

			for ix, r := range mask {
				if r == '1' {
					value = setBit(value, uint(ix))
				}
				if r == '0' {
					value = clearBit(value, uint(ix))
				}
			}
			mem[addr] = value
			continue
		}
		return 0, fmt.Errorf("Line did not match Res [%s]", line)
	}
	var res uint64
	for _, v := range mem {
		res += v
	}
	return res, nil
}

func day14Part2(input []string) (uint64, error) {

	maskRe := regexp.MustCompile(`^mask = ([X10]+)$`)
	memRe := regexp.MustCompile(`^mem\[([0-9]+)\] = ([0-9]+)$`)

	setBit := func(n uint64, pos uint) uint64 {
		n |= (1 << pos)
		return n
	}
	clearBit := func(n uint64, pos uint) uint64 {
		mask := ^(1 << pos)
		n &= uint64(mask)
		return n
	}
	reverse := func(s string) string {
		rns := []rune(s) // convert to rune
		for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

			// swap the letters of the string,
			// like first with last and so on.
			rns[i], rns[j] = rns[j], rns[i]
		}
		// return the reversed string.
		return string(rns)
	}

	mem := make(map[uint64]uint64)
	mask := ""
	for _, line := range input {

		// either the line will be a pattern

		// or the line will be an assignment
		match := maskRe.FindStringSubmatch(line)
		if match != nil {
			// we have a new mask
			mask = reverse(match[1])
			continue
		}
		match = memRe.FindStringSubmatch(line)
		if match != nil {
			// we have a new memory line
			addr, err := strconv.ParseUint(match[1], 10, 32)
			if err != nil {
				return 0, err
			}
			value, err := strconv.ParseUint(match[2], 10, 64)
			if err != nil {
				return 0, err
			}

			addrs := make([]uint64, 1)
			addrs[0] = uint64(addr)

			for ix, r := range mask {
				switch r {
				case '1':
					for aix, v := range addrs {
						addrs[aix] = setBit(v, uint(ix))
					}
				case 'X':
					for aix, v := range addrs {
						addrs[aix] = setBit(v, uint(ix))
						addrs = append(addrs, clearBit(v, uint(ix)))
					}
				}
			}
			for _, addr := range addrs {
				mem[addr] = value
			}
			continue
		}
		return 0, fmt.Errorf("Line did not match Res [%s]", line)
	}
	var res uint64
	for _, v := range mem {
		res += v
	}
	return res, nil
}
