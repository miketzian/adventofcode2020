package adventofcode2020

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func day21(input []string) (int, error) {

	agRegex, err := regexp.Compile(` \(contains ([^\)]+)\)`)
	if err != nil {
		return 0, err
	}

	foodcounts := make(map[string]int)
	maybeAllergens := make(map[string]map[string]bool)

	for _, line := range input {

		// replace all the instances where there are allergens
		m := agRegex.FindStringSubmatchIndex(line)
		if len(m) == 0 {
			return 0, fmt.Errorf("no allergens?")
		}
		agList := strings.Split(line[m[2]:m[3]], ", ")
		line = line[0:m[0]]

		for _, food := range strings.Split(line, " ") {
			foodcounts[food]++
		}

		for _, ag := range agList {

			thisAg := make(map[string]bool)
			for _, food := range strings.Split(line, " ") {
				thisAg[food] = true
			}

			if _, prs := maybeAllergens[ag]; !prs {
				// add all of them
				maybeAllergens[ag] = thisAg
			} else {
				// only foods that are allergens in both should stay
				for k := range maybeAllergens[ag] {
					if _, prs := thisAg[k]; !prs {
						delete(maybeAllergens[ag], k)
					}
				}
			}
		}
	}

	isAllergen := make(map[string]string)
	hasAllergen := make(map[string]string)
	loops := 0
	for len(maybeAllergens) > 0 && loops < 50 {
		// find a maybe allergen that has only one option
		for allergen, v := range maybeAllergens {
			if len(v) == 1 {
				// we will stipulate that this is the allergen
				// and remove it from all others
				delete(maybeAllergens, allergen)
				for food := range v {
					isAllergen[food] = allergen
					hasAllergen[allergen] = food

					// also remove where this food exists in other lists
					// as the allergen is now identified
					for allKey, otherMaybe := range maybeAllergens {
						if _, prs := otherMaybe[food]; prs {
							delete(otherMaybe, food)
						}
						// this shouldn't occur as then the other allergen is unidentified
						if len(otherMaybe) == 0 {
							return 0, fmt.Errorf("Unable to identify [%s]", allKey)
						}
					}
				}
				break
			}
		}
		loops++
	}
	if loops == 50 {
		return 0, fmt.Errorf("too many loops")
	}

	safetotal := 0
	for code, count := range foodcounts {
		if _, prs := isAllergen[code]; !prs {
			// fmt.Printf("%s does not contain allergen adding %d\n", code, count)
			safetotal += count
			// } else {
			//		fmt.Printf("%s contains allergens, skipping %d\n", code, count)
		}
	}

	keys := make([]string, 0, len(hasAllergen))
	for k := range hasAllergen {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	danger := ""
	for _, k := range keys {
		if len(danger) != 0 {
			danger = fmt.Sprintf("%s,%s", danger, hasAllergen[k])
		} else {
			danger = hasAllergen[k]
		}
	}
	fmt.Printf("%s\n", danger)
	return safetotal, nil
}
