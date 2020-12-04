package adventofcode2020

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type dayFourValidity struct {
	colorRegex *regexp.Regexp
}

func dayFourValidation() *dayFourValidity {
	return &dayFourValidity{colorRegex: regexp.MustCompile("^#[a-z0-9]+$")}
}

func (o *dayFourValidity) validNum(value string, digits int, min int64, max int64) bool {
	if len(value) != digits {
		return false
	}
	intValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return false
	}
	if intValue < min {
		return false
	}
	if intValue > max {
		return false
	}
	return true
}

func (o *dayFourValidity) validSuffix(value string, suffix string, min int64, max int64) bool {
	if !strings.HasSuffix(value, suffix) {
		return false
	}
	intValue, err := strconv.ParseInt(value[:len(value)-2], 10, 32)
	if err != nil {
		return false
	}
	if intValue < min {
		return false
	}
	if intValue > max {
		return false
	}
	return true
}

// func (o *dayFourValidity) validEye(value) {

// }

func dayFour(input []string, extraValidation bool) (int, error) {
	// passports valid
	validityChecker := dayFourValidation()

	valid := 0
	checkValid := func(input string) bool {

		// byr (Birth Year)
		// iyr (Issue Year)
		// eyr (Expiration Year)
		// hgt (Height)
		// hcl (Hair Color)
		// ecl (Eye Color)
		// pid (Passport ID)
		// cid (Country ID)

		required := []string{
			"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
		}
		// skip cid

		data := make(map[string]string, 0)
		for _, v := range strings.Split(input, " ") {
			kv := strings.Split(v, ":")
			data[kv[0]] = kv[1]
		}
		for _, k := range required {
			if v, exists := data[k]; !exists {
				return false
			} else if extraValidation {
				switch {
				case k == "byr":
					if !validityChecker.validNum(v, 4, 1920, 2002) {
						return false
					}
				case k == "iyr":
					if !validityChecker.validNum(v, 4, 2010, 2020) {
						return false
					}
				case k == "eyr":
					if !validityChecker.validNum(v, 4, 2020, 2030) {
						return false
					}
				case k == "hgt":
					if len(v) < 3 {
						return false
					}
					if !validityChecker.validSuffix(v, "in", 59, 76) && !validityChecker.validSuffix(v, "cm", 150, 193) {
						return false
					}
				case k == "hcl":
					// (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
					if len(v) != 7 || !validityChecker.colorRegex.MatchString(v) {
						return false
					}
				case k == "ecl":
					// (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
					validEye := false
					for _, eye := range strings.Split("amb blu brn gry grn hzl oth", " ") {
						if v == eye {
							validEye = true
							break
						}
					}
					if !validEye {
						return false
					}
				case k == "pid":
					// pid (Passport ID) - a nine-digit number, including leading zeroes.
					if !validityChecker.validNum(v, 9, 0, 999999999) {
						return false
					}
				}
			}
		}
		return true
	}
	curr := ""
	for _, v := range input {
		v = strings.Trim(v, " \t\r")
		if v == "" {
			if checkValid(curr) {
				valid++
			}
			curr = ""
		} else if curr == "" {
			curr = v
		} else {
			curr = fmt.Sprintf("%s %s", curr, v)
		}
	}
	if curr != "" {
		if checkValid(curr) {
			valid++
		}
	}
	return valid, nil
}
