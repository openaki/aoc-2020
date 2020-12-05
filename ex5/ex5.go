package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func GetRowId(code string) int {
	rowNum := 0
	for _, c := range code[:len(code) - 3] {
		rowNum = rowNum << 1
		if c == 'B' {
			rowNum += 1
		}
	}

	colNum := 0
	for _, r := range code[len(code) - 3:] {
		colNum = colNum << 1
		if r == 'R' {
			colNum += 1
		}
	}
	seatId := rowNum * 8 + colNum
	return seatId
}

func GetContent() []string {
	content, err := ioutil.ReadFile("./input5.txt")
	if err != nil {
		log.Fatal(err)
	}

	seatIds := strings.Split(string(content), "\n")

	return seatIds;
}

func Solve() {
	seatIds := GetContent()
	maxId := 0
	seatsSeen := make(map[int]bool)
	for _, s := range seatIds {
		if s == "" {
			continue
		}
		id := GetRowId(s)
		seatsSeen[id] = true
		if id > maxId {
			maxId = id
		}
	}
	fmt.Println("Solution Ex 5 part a ", maxId)

	for i := 0; i < maxId; i++ {

		b := seatsSeen[i]
		if !b {
			if seatsSeen[i+1] && seatsSeen[i-1] {
				fmt.Println("Solution Ex 5 part b", i)

			}
		}
	}
}

func main() {
	Solve()
}
