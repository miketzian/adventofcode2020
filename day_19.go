package adventofcode2020

import (
	"fmt"
	"strconv"
	"strings"
)

type day19data struct {
	rules map[int]*day19Rule
}

func (obj *day19data) matches(ruleIx int, dataOpts []string) (bool, []string) {

	rule := obj.rules[ruleIx]

	var newOpts []string

	for _, data := range dataOpts {
		if len(data) == 0 {
			// no more data to match in this path
			continue
		}
		if rule.char != "" {
			if data[0:1] == rule.char {
				newOpts = append(newOpts, data[1:])
			}
		} else {
		ruleLinks:
			for _, cond := range rule.links {
				matches := true
				matchData := []string{data}
				for _, subRuleIx := range cond {
					matches, matchData = obj.matches(subRuleIx, matchData)
					if !matches {
						// try the next rule
						continue ruleLinks
					}
				}
				newOpts = append(newOpts, matchData...)
			}
		}
	}
	return newOpts != nil, newOpts
}

type day19Rule struct {
	char  string
	links [][]int
}

func day19(input []string, part2 bool) (int, error) {

	rules, data, err := day19Load(input, part2)
	if err != nil {
		return 0, err
	}

	obj := &day19data{rules}
	matches := 0

lineIter:
	for _, line := range data {
		// determine if rule 0 matches
		if match, rem := obj.matches(0, []string{line}); match {
			for _, r := range rem {
				if len(r) == 0 {
					// fmt.Printf("Line Matched: [%s]\n", line)
					matches++
					continue lineIter
				}
			}
		}
	}
	return matches, nil
}

func day19Load(input []string, part2 bool) (map[int]*day19Rule, []string, error) {

	data := make(map[int]*day19Rule)

	for ix, line := range input {
		if line == "" {
			return data, input[ix+1:], nil
		}
		kv := strings.Split(line, ": ")

		key, err := strconv.Atoi(kv[0])
		if err != nil {
			return nil, nil, err
		}

		if kv[1][0] == '"' {
			// we have a character
			data[key] = &day19Rule{kv[1][1:2], nil}
		} else {
			// we need to parse the or condition
			data[key] = &day19Rule{"", make([][]int, 0)}

			if part2 {
				switch key {
				case 8:
					kv[1] = "42 | 42 8"
				case 11:
					kv[1] = "42 31 | 42 11 31"
				}
			}

			next := make([]int, 0)
			for _, v := range strings.Split(kv[1], " ") {
				if v == "|" {
					data[key].links = append(data[key].links, next)
					next = make([]int, 0)
				} else {
					value, err := strconv.Atoi(v)
					if err != nil {
						return nil, nil, err
					}
					next = append(next, value)
				}
			}
			data[key].links = append(data[key].links, next)

		}
	}
	return nil, nil, fmt.Errorf("The gap wasn't encountered")
}
