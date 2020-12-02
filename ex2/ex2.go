package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ValidatePasswordsA(inputs []string) int {

	re := regexp.MustCompile("(\\d+)-(\\d+) ([^:]*): ([^ ]*)")

	acceptedCount := 0
	for _, input := range inputs {
		matchString := re.FindStringSubmatch(input)
		lowerCount, err := strconv.Atoi(matchString[1])

		if err != nil {
			panic("lower count failed to parse")
		}
		upperCount, err := strconv.Atoi(matchString[2])
		if err != nil {
			panic("upper count failed to parse")
		}

		pattern, password := matchString[3], matchString[4]

		count := strings.Count(password, pattern)

		accepted := count >= lowerCount && count <= upperCount
		if accepted {
			acceptedCount++
		}
	}

	return acceptedCount
}
func ValidatePasswordsB(inputs []string) int {

	re := regexp.MustCompile("(\\d+)-(\\d+) ([^:]*): ([^ ]*)")

	acceptedCount := 0
	for _, input := range inputs {
		matchString := re.FindStringSubmatch(input)
		index1, err := strconv.Atoi(matchString[1])

		if err != nil {
			panic("lower count failed to parse")
		}
		index2, err := strconv.Atoi(matchString[2])
		if err != nil {
			panic("upper count failed to parse")
		}

		pattern, password := matchString[3][0], matchString[4]

		matchFound := 0
		if password[index1 - 1] == pattern {
			matchFound++
		}
		if password[index2 - 1] == pattern {
			matchFound++
		}
		if matchFound == 1{
			acceptedCount++
		}
	}

	return acceptedCount
}

func main() {

	data, err := ioutil.ReadFile("./input2.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	inputStr := strings.TrimSpace(string(data))
	inputs := strings.Split(inputStr, "\n")
	solvea := ValidatePasswordsA(inputs)
	fmt.Printf("Solution for part a: %d\n", solvea)
	solveb := ValidatePasswordsB(inputs)
	fmt.Printf("Solution for part b: %d\n", solveb)
}
