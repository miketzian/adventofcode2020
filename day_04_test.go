package adventofcode2020

import (
	"testing"
)

func TestDayFour(t *testing.T) {

	testData := []string{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"	",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"	",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in"}

	testInvalid := []string{
		"eyr:1972 cid:100",
		"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
		"",
		"iyr:2019",
		"hcl:#602927 eyr:1967 hgt:170cm",
		"ecl:grn pid:012533040 byr:1946",
		"",
		"hcl:dab227 iyr:2012",
		"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
		"",
		"hgt:59cm ecl:zzz",
		"eyr:2038 hcl:74454a iyr:2023",
		"pid:3556412378 byr:2007",
	}

	testValid := []string{
		"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
		"hcl:#623a2f",
		"",
		"eyr:2029 ecl:blu cid:129 byr:1989",
		"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
		"",
		"hcl:#888785",
		"hgt:164cm byr:2001 iyr:2015 cid:88",
		"pid:545766238 ecl:hzl",
		"eyr:2022",
		"",
		"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
	}

	inputData, err := readFileAsStrings("day_04_input.txt")
	if err != nil {
		t.Error(err)
	}
	if len(inputData) == 0 || len(testInvalid) == 0 || len(testValid) == 0 || len(testData) == 0 {
		t.Fatalf("check")
	}
	cases := []struct {
		input           []string
		testcase        string
		extraValidation bool
		expected        int
	}{
		{testData, "unit", false, 2},
		{inputData, "first", false, -1},
		{testInvalid, "test invalid", true, 0},
		{testValid, "test valid", true, 4},
		{inputData, "second", true, -1},
	}
	// mult := 1
	for _, c := range cases {
		result, err := dayFour(c.input, c.extraValidation)
		if err != nil {
			t.Fatal(err)
		}
		if c.expected == -1 {
			// we don't know the answer, we're using
			// the computation
			t.Logf("Result (%s): %d\n", c.testcase, result)
			continue
		}
		if result != c.expected {
			t.Errorf("Result (%s) %d != Expected %d",
				c.testcase, result, c.expected)
		}
	}
	// t.Logf("Result: %d\n", mult)
}

func TestDayFourValidation(t *testing.T) {

	validityChecker := dayFourValidation()

	if !validityChecker.validNum("1980", 4, 1970, 2000) {
		t.Errorf("Valid Num Failed")
	}
	if !validityChecker.validNum("1980", 4, 1980, 2000) {
		t.Errorf("Valid Num Failed")
	}
	if !validityChecker.validNum("2000", 4, 1980, 2000) {
		t.Errorf("Valid Num Failed")
	}
	if validityChecker.validNum("1970", 4, 1980, 2000) {
		t.Errorf("Valid Num Failed")
	}
	if validityChecker.validNum("2001", 4, 1980, 2000) {
		t.Errorf("Valid Num Failed")
	}
	if !validityChecker.validSuffix("150cm", "cm", 149, 170) {
		t.Errorf("Valid suffix Failed")
	}
	if validityChecker.validSuffix("150cm", "cm", 151, 170) {
		t.Errorf("Valid suffix Failed")
	}
	if validityChecker.validSuffix("150cm", "cm", 120, 149) {
		t.Errorf("Valid suffix Failed")
	}
}
