package adventofcode2020

import (
	"regexp"
	"strconv"
)

// BagDag - bag instance plus what it can contain
type BagDag struct {
	bagColor       string
	containedIn    []*BagDag
	contains       []*BagDag
	containedCount []int
}

func (bag *BagDag) bagsIn() int {
	total := 1
	for ix, bagContains := range bag.contains {
		cnt := bag.containedCount[ix]
		// fmt.Printf("Bag [%s] Contains: %d - '%s'\n",
		// 	bag.bagColor, cnt, bagContains.bagColor)
		total += (cnt * bagContains.bagsIn())
	}
	// fmt.Printf("\t\tBag [%s] Contains Total: %d\n", bag.bagColor, total)
	return total
}

func daySeven(input []string) (int, int, error) {

	bagRe := regexp.MustCompile(`^(\w+ \w+) bags contain`)
	// we're ignoring the # of bags for now
	bagContainsRe := regexp.MustCompile(`([0-9]+) (\w+ \w+) bag`)

	allBags := make(map[string]*BagDag)

	getOrMakeBag := func(color string) *BagDag {
		if _, present := allBags[color]; !present {
			allBags[color] = &BagDag{
				bagColor:       color,
				containedIn:    make([]*BagDag, 0),
				contains:       make([]*BagDag, 0),
				containedCount: make([]int, 0),
			}
		}
		return allBags[color]
	}

	for _, v := range input {
		bagColor := bagRe.FindStringSubmatch(v)[1]
		// fmt.Printf("Input [%s]\nBag: %s\n", v, bagColor)
		bag := getOrMakeBag(bagColor)

		for _, v := range bagContainsRe.FindAllStringSubmatch(v, -1) {
			bagContains := getOrMakeBag(v[2])
			bagContains.containedIn = append(bagContains.containedIn, bag)

			// reverse, for counts
			bag.contains = append(bag.contains, bagContains)
			cnt, err := strconv.ParseInt(v[1], 10, 64)
			if err != nil {
				return -1, -1, err
			}
			bag.containedCount = append(bag.containedCount, int(cnt))
		}
	}

	count := make(map[string]bool)
	var bagCheck func(bag *BagDag)
	bagCheck = func(bag *BagDag) {
		count[bag.bagColor] = true
		for _, inBag := range bag.containedIn {
			if _, prs := count[inBag.bagColor]; !prs {
				// not present -> not yet traversed (there may be more than one path)
				bagCheck(inBag)
			}
		}
	}
	shinyGold := allBags["shiny gold"]
	bagCheck(shinyGold)

	// minus one to remove shiny gold itself
	return len(count) - 1, shinyGold.bagsIn() - 1, nil
}
