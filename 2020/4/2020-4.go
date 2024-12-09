package main

import (
	. "github.com/bendiscz/aoc/aoc"
	"regexp"
	"strconv"
)

const (
	year    = 2020
	day     = 4
	example = `

ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in

`
)

func main() {
	Run(year, day, example, solve)
}

func solve(p *Problem) {
	s1, s2 := 0, 0
	for {
		pass := readPassport(p)
		if len(pass) == 0 {
			break
		}

		if len(pass) == 7 {
			s1++
			if validate(pass) {
				s2++
			}
		}
	}

	p.PartOne(s1)
	p.PartTwo(s2)
}

func readPassport(p *Problem) map[string]string {
	pass := map[string]string{}
	for p.NextLine() {
		if p.Line() == "" {
			break
		}

		f := SplitFieldsDelim(p.Line(), " :")
		for i := 0; i < len(f); i += 2 {
			pass[f[i]] = f[i+1]
		}
	}
	delete(pass, "cid")
	return pass
}

var heightRE = regexp.MustCompile(`^(\d{2,3})(in|cm)$`)
var hairRE = regexp.MustCompile(`^#([0-9a-f]{6})$`)
var eyeRE = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
var pidRE = regexp.MustCompile(`^(\d{9})$`)

func validate(pass map[string]string) bool {
	return validateYear(pass["byr"], 1920, 2002) &&
		validateYear(pass["iyr"], 2010, 2020) &&
		validateYear(pass["eyr"], 2020, 2030) &&
		validateHeight(pass["hgt"]) &&
		hairRE.MatchString(pass["hcl"]) &&
		eyeRE.MatchString(pass["ecl"]) &&
		pidRE.MatchString(pass["pid"])
}

func validateYear(s string, a, b int) bool {
	x, _ := strconv.Atoi(s)
	return x >= a && x <= b
}

func validateHeight(s string) bool {
	m := heightRE.FindStringSubmatch(s)
	if len(m) != 3 {
		return false
	}
	x := ParseInt(m[1])
	return m[2] == "in" && x >= 59 && x <= 76 || m[2] == "cm" && x >= 150 && x <= 193
}
