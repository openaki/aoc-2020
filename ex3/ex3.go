package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)


func CountTrees(topology string, downDX int, rightDX int) int {

	trimmed := strings.TrimSpace(topology)
	mapVec := strings.Split(trimmed, "\n")

	endCol := len(mapVec)

	cycle := len(mapVec[0])

	treeCount := 0

	fmt.Println(downDX, rightDX)
	for row, col := downDX, 0; row < endCol; row += downDX {

		col += rightDX

		if row == endCol {
			break
		}

		if mapVec[row][col%cycle] == '#' {
			treeCount++
		}
	}

	return treeCount
}

func CountTreesDifferentSlopes(topology string) int {
	slopes := []struct {
		r, d int
	}{
		{r: 1, d: 1},
		{r: 3, d: 1},
		{r: 5, d: 1},
		{r: 7, d: 1},
		{r: 1, d: 2},
	}

	solveb := 1

	for _, v := range slopes {
		solveb *= CountTrees(topology, v.d, v.r)

	}

	return solveb
}

func main() {
	content, err := ioutil.ReadFile("./input3.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)

	}

	contentStr := string(content)
	solvea := CountTrees(contentStr, 1, 3)

	fmt.Printf("Solution for 3A: %d\n", solvea)

	solveb := CountTreesDifferentSlopes(contentStr)

	fmt.Printf("Solution for 3B: %d\n", solveb)

}
