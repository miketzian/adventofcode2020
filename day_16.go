package adventofcode2020

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	field  string
	ranges [][]int
}

func (r *rule) matches(value int) bool {
	for _, r := range r.ranges {
		if value >= r[0] && value <= r[1] {
			return true
		}
	}
	return false
}

type ticket struct {
	data []int
}

func (t *ticket) matchErrors(rules []*rule) int {
	errors := 0
outer:
	for _, d := range t.data {
		for _, r := range rules {
			if r.matches(d) {
				// field matches at least one rule
				continue outer
			}
		}
		// this field matches no rules
		// return false
		errors += d
	}
	return errors
}

// Day16Data ...
type Day16Data struct {
	fieldRules    []*rule
	yourTicket    []int
	nearbyTickets []*ticket
}

func day16Load(input []string) (*Day16Data, error) {

	section := 0
	ruleRegexp, err := regexp.Compile(`([0-9]+)\-([0-9]+)`)
	if err != nil {
		return nil, err
	}

	strArrToInt := func(input []string) ([]int, error) {

		res := make([]int, len(input))
		for ix, v := range input {
			i, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			res[ix] = i
		}
		return res, nil
	}

	var fieldRules []*rule
	var yourTicket []int
	var nearbyTickets []*ticket

	for _, line := range input {
		// fmt.Printf("Line: [%s], Ix [%d], Section [%d]\n",
		// 	line, lix, section)
		if line == "" {
			section++
			continue
		}
		switch section {
		case 0:
			// parse key=value
			kv := strings.Split(line, ": ")
			// fmt.Printf("Line: [%s], Field [%s], Ranges [%s]\n",
			// 	line, kv[0], kv[1])
			field := kv[0]
			fieldRule := &rule{field, [][]int{}}
			// [][]string -> [][]int
			for _, v := range ruleRegexp.FindAllStringSubmatch(kv[1], -1) {
				fieldRange, err := strArrToInt(v[1:])
				if err != nil {
					return nil, err
				}
				fieldRule.ranges = append(fieldRule.ranges, fieldRange)
			}
			fieldRules = append(fieldRules, fieldRule)
		case 1:
			if line == "your ticket:" {
				continue
			}
			if yourTicket != nil {
				return nil, fmt.Errorf("Somehow your ticket already defined?")
			}
			yourTicket, err = strArrToInt(strings.Split(line, ","))
			if err != nil {
				return nil, err
			}
		case 2:
			if line == "nearby tickets:" {
				continue
			}
			nearbyTicket, err := strArrToInt(strings.Split(line, ","))
			if err != nil {
				return nil, err
			}
			nearbyTickets = append(nearbyTickets, &ticket{data: nearbyTicket})
		default:
			return nil, fmt.Errorf("Unexpected section [%d]", section)
		}
	}

	fmt.Printf("Loaded %d rules, %v yours, %d nearby",
		len(fieldRules), yourTicket != nil, len(nearbyTickets))

	return &Day16Data{fieldRules, yourTicket, nearbyTickets}, nil
}

func (load *Day16Data) yourTicketData() (map[string]int, error) {

	fieldRules := load.fieldRules
	yourTicket := load.yourTicket
	nearbyTickets := load.nearbyTickets
	validTickets := make([]*ticket, 1)
	validTickets[0] = &ticket{yourTicket}
	for _, t := range nearbyTickets {
		ticketValid := true
		for _, v := range t.data {
			fieldValid := false
			for _, r := range fieldRules {
				if r.matches(v) {
					fieldValid = true
					break
				}
			}
			if !fieldValid {
				ticketValid = false
			}
			if !ticketValid {
				break
			}
		}
		if ticketValid {
			validTickets = append(validTickets, t)
		}
	}

	if validTickets[0].matchErrors(fieldRules) != 0 {
		return nil, fmt.Errorf("We got errors for our ticket %v", yourTicket)
	}

	// for those valid tickets

	fieldOk := make(map[string]int)
	fieldMatched := make(map[int]bool)
	for _, rule := range fieldRules {
		fieldOk[rule.field] = -1
	}
	allOk := func() bool {
		for _, ok := range fieldOk {
			if ok == -1 {
				return false
			}
		}
		return true
	}
	isOk := func(rule *rule) bool {
		return fieldOk[rule.field] != -1
	}
	loops := 0
	maxLoops := 20
	ticketFields := len(yourTicket)
	for !allOk() && loops < maxLoops {
		fmt.Println("** loop **")

		// look for fields for which there is only one possible field
	nextRule:
		for _, fieldRule := range fieldRules {
			if isOk(fieldRule) {
				continue
			}
			ticketField := 0
			ruleTicketField := -1
		nextField:
			for ticketField < ticketFields {
				if fieldMatched[ticketField] {
					// we already know the field for this index
					ticketField++
					continue
				}
				for _, validTicket := range validTickets {
					if !fieldRule.matches(validTicket.data[ticketField]) {
						ticketField++
						continue nextField
					}
				}
				// if we get here then this rule matches all tickets
				// for index ticketField
				if ruleTicketField == -1 {
					// this is the first that matched
					ruleTicketField = ticketField
				} else {
					// this rule matches multiple remaining fields
					// lets skip for now.
					// fmt.Printf("field=[%s] could be %d or %d\n", fieldRule.field, ticketField, ruleTicketField)
					continue nextRule
				}
				ticketField++
			}
			// this rule matched one and only one field
			if ruleTicketField != -1 {
				fmt.Printf("field=[%s] is %d\n", fieldRule.field, ruleTicketField)
				fieldOk[fieldRule.field] = ruleTicketField
				fieldMatched[ruleTicketField] = true
			}
		}
		loops++
	}
	if loops == maxLoops {
		return nil, fmt.Errorf("max loops reached ok=%v", fieldOk)
	}

	result := make(map[string]int)
	for _, fieldRule := range fieldRules {
		result[fieldRule.field] = yourTicket[fieldOk[fieldRule.field]]
	}
	return result, nil
}

func day16(input []string) (int, error) {

	load, err := day16Load(input)
	if err != nil {
		return 0, err
	}
	fieldRules := load.fieldRules
	yourTicket := load.yourTicket
	nearbyTickets := load.nearbyTickets

	fmt.Printf("%v\n", yourTicket)
	ticketErrors := 0
	for _, t := range nearbyTickets {
		ticketErrors += t.matchErrors(fieldRules)
	}
	return ticketErrors, nil
}

func day16MapFields(input []string) (map[string]int, error) {
	load, err := day16Load(input)
	if err != nil {
		return nil, err
	}
	return load.yourTicketData()
}

func day16Part2(input []string) (int, error) {

	load, err := day16Load(input)
	if err != nil {
		return 0, err
	}

	ticket, err := load.yourTicketData()
	if err != nil {
		return 0, err
	}
	result := 1
	for key, value := range ticket {
		if strings.HasPrefix(key, "departure") {
			result *= value
		}

	}
	return result, nil
}
