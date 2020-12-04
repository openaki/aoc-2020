package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var validKeys map[string]bool = map[string]bool{
	"byr": true,
	"iyr": true,
	"eyr": true,
	"hgt": true,
	"hcl": true,
	"ecl": true,
	"pid": true,
	"cid": true,
}


func ValidateRecords(records map[string]string) bool {
	validEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	validEyeColorMap := make(map[string]bool)
	for _, v := range validEyeColors {
		validEyeColorMap[v] = true
	}

	validateNum := func(s string, low, high int) bool {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		if low <= num && num <= high {
			return true
		}
		return false
	}

	hclRe, err := regexp.Compile(`^#[0-9a-f]{6}$`)
	if err != nil {
		return false
	}
	pidRe, err := regexp.Compile(`^\d{9}$`)
	if err != nil {
		return false
	}

	for k, v := range records {
		switch k {
		case "byr":
			{
				if !validateNum(v, 1920, 2002) {
					return false
				}
			}
		case "iyr":
			{
				if !validateNum(v, 2010, 2020) {
					return false
				}
			}
		case "eyr":
			{
				if !validateNum(v, 2020, 2030) {
					return false
				}
			}
		case "hgt":
			{
				hnum := v[:len(v)-2]
				if strings.HasSuffix(v, "cm") && !validateNum(hnum, 150, 193) {
					return false
				} else if strings.HasSuffix(v, "in") && !validateNum(hnum, 59, 76) {
					return false
				} else if !strings.HasSuffix(v, "cm") && !strings.HasSuffix(v, "in") {
					return false
				}
			}
		case "hcl":
			{

				if !hclRe.MatchString(v) {
					return false
				}
			}
		case "ecl":
			{
				_, exists := validEyeColorMap[v]
				if !exists {
					return false
				}

			}
		case "pid":
			{
				regexMatch := true
				if !pidRe.MatchString(v) {
					regexMatch = false
					return false
				}

				_ = regexMatch
			}
		}
	}
	return true

}

func CountValidPassorts(batch string, validateRecords bool) int {

	lines := strings.Split(batch, "\n")

	currentRecords := make(map[string]string)
	invalidRecordCount := 0
	validRecordCount := 0

	endOfRecord := func() {
		_, cidExists := currentRecords["cid"]

		keysValid := !validateRecords || ValidateRecords(currentRecords)

		if len(currentRecords) == 8 && keysValid {
			validRecordCount++
		} else if len(currentRecords) == 7 && !cidExists && keysValid {
			validRecordCount++
		} else {
			invalidRecordCount++
		}

		//fmt.Println(currentRecords, validRecordCount)
		currentRecords = make(map[string]string)

	}
	for _, l := range lines {

		l = strings.TrimSpace(l)
		words := strings.Split(l, " ")

		if len(words) == 1 && words[0] == "" {
			endOfRecord()
			continue
		}

		for _, w := range words {
			kv := strings.Split(w, ":")
			key := strings.TrimSpace(kv[0])
			var value = strings.TrimSpace(kv[1])
			_, exists := validKeys[key]
			if exists {
				currentRecords[key] = value
			}
		}

	}
	endOfRecord()
	return validRecordCount
}

func main() {
	contents, err := ioutil.ReadFile("./input4.txt")
	if err != nil {
		log.Fatal(err)
	}

	solvea := CountValidPassorts(string(contents), false)
	fmt.Printf("Soulution Ex 4 part a: %d\n", solvea)

	solveb := CountValidPassorts(string(contents), true)
	fmt.Printf("Soulution Ex 4 part b: %d\n", solveb)
}
