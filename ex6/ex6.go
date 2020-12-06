package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Mode int

const (
	Union Mode = iota
	Intersection
)

func GetTotalYesAnswers(content string, mode Mode) int {

	lines := strings.Split(content, "\n")

	answersYes := make(map[int]int)

	numOfPeopleInGroup := 0
	sumOfYes := 0
	endOfGroup := func() {

		switch mode {
		case Intersection:
			for _, ys := range answersYes{
				if ys == numOfPeopleInGroup {
					sumOfYes++
				}
		}
		case Union :
			sumOfYes += len(answersYes)
		}
		answersYes = make(map[int]int)
		numOfPeopleInGroup = 0
	}

	for _, l := range lines {
		individualAnswers := make(map[int]bool)
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			endOfGroup()
			continue
		}
		numOfPeopleInGroup++

		for _, a := range l {
			index := int(a)

			if individualAnswers[index] {
				continue
			}
			individualAnswers[index] = true
			answersYes[index] += 1
		}
	}

	endOfGroup()

	return sumOfYes
}

func main() {
	content, err := ioutil.ReadFile("./input6.txt")
	if err != nil {
		log.Fatal(err)
	}
	answer := GetTotalYesAnswers(string(content), Union)

	fmt.Println("Solution ex5 part a ", answer)

	answer = GetTotalYesAnswers(string(content), Intersection)

	fmt.Println("Solution ex5 part b ", answer)
}
