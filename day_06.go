package adventofcode2020

func daySix(input []string) (int, error) {

	group := make(map[rune]bool)
	total := 0
	for _, v := range input {
		if v == "" {
			// group is finished
			total += len(group)
			group = make(map[rune]bool)
			continue
		}
		for _, q := range v {
			if _, r := group[q]; !r {
				group[q] = true
			}
		}
	}
	// add the last group
	return total + len(group), nil
}

func daySixEveryone(input []string) (int, error) {

	group := make(map[rune]int)
	groupSize := 0
	total := 0
	for _, v := range input {
		if v == "" {
			// group is finished
			for _, answeredYes := range group {
				if answeredYes == groupSize {
					// everyone answered yes
					total++
				}
			}
			group = make(map[rune]int)
			groupSize = 0
			continue
		}
		for _, q := range v {
			if _, r := group[q]; !r {
				group[q] = 1
			} else {
				group[q]++
			}
		}
		groupSize++
	}
	for _, answeredYes := range group {
		if answeredYes == groupSize {
			// everyone answered yes
			total++
		}
	}
	// add the last group
	return total, nil
}
