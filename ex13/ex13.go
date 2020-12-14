package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func GetNextAvaiableBus(content string) int {
	lines := strings.Split(content, "\n")
	depart, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatal(err)
	}
	options := strings.Split(lines[1], ",")
	earliest := math.MaxInt32
	busId := -1

	for _, o := range options {
		if o == "x" {
			continue
		}
		freq, err := strconv.Atoi(o)
		if err != nil {
			log.Fatal(err)
		}
		waitingTime := freq - (depart % freq)

		if waitingTime < earliest {
			earliest = waitingTime
			busId = freq
		}
		fmt.Println(waitingTime, depart, busId, waitingTime, earliest)
	}

	return earliest * busId
}

func BruteForce(content string) int {
	lines := strings.Split(content, "\n")
	options := strings.Split(lines[1], ",")

	requirements := make([]int, 0)
	maxFreq := -1

	var first int = -1
	for _, o := range options {
		if o == "x" {
			requirements = append(requirements, 0)
			continue
		}
		freq, err := strconv.Atoi(o)
		if err != nil {
			log.Fatal(err)
		}
		if first == -1{
			first = freq
		}
		requirements = append(requirements, freq)
		if maxFreq < freq {
			maxFreq = freq
		}
	}
	fmt.Println(requirements)

	step := 1
	ans := 0
	for delay, id := range requirements {
		if id == 0 {
			continue
		}
		for i := ans; true; i += step {
			if (i + delay) % id == 0 {
				ans = i
				step *= id // assumes only primes
				break
			}
		}
	}

	return ans
}

func main() {
	content, err := ioutil.ReadFile("./input13.txt")
	if err != nil {
		log.Fatal(err)
	}

	ans := GetNextAvaiableBus(string(content))
	fmt.Println(ans)

	ans = BruteForce(string(content))
	fmt.Println(ans)
}
