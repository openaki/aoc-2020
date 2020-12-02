package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Solve1a(input []int) int {

	index1 := -1
	index2 := -1

	for i, a := range input {
		for _, b := range input[i:] {
			if a+b == 2020 {
				index1 = a
				index2 = b
			}

		}
	}

	return index1 * index2

}

func Solve1b(input []int) int {
	index1 := -1
	index2 := -1
	index3 := -1

	for i, a := range input {
		for _, b := range input[i:] {
			for _, c := range input[i:] {
				if a+b+c == 2020 {
					index1 = a
					index2 = b
					index3 = c
				}
			}

		}
	}

	return index1 * index2 * index3
}

func main() {

	data, err := ioutil.ReadFile("./inputs/input1.txt")
	if err != nil {
		log.Fatalf("Failed to open file %v", err)
		return
	}
	lines := string(data)
	lines = strings.TrimSpace(lines)
	lineVec := strings.Split(lines, "\n")

	inputs := make([]int, len(lineVec))
	for i, v := range lineVec {
		inputs[i], err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("failed to prase number %d %s: %v", i, v, err)
			return
		}
	}

	ans := Solve1a(inputs)
	fmt.Printf("Solution 1a: %d\n", ans)

	ans = Solve1b(inputs)
	fmt.Printf("Solution 1b: %d\n", ans)

}
