package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func handleDirection(dir int, val int, x, y int) (int, int) {
	switch dir {
	case 'N':
		{
			y += val
		}
	case 'S':
		{
			y -= val
		}
	case 'E':
		{
			x += val
		}
	case 'W':
		{
			x -= val
		}
	}
	return x, y
}

func RunInstructions(content string) int {

	instructions := strings.Split(content, "\n")

	x, y := 0, 0
	direction := 0

	dirArray := []int{'E', 'S', 'W', 'N'}

	for _, instr := range instructions {
		cmd := instr[0]
		val, err := strconv.Atoi(instr[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch cmd {
		case 'F':
			{
				x, y = handleDirection(dirArray[direction], val, x, y)
			}
		case 'R':
			{
				moves := val / 90
				moves = moves % 4
				direction = (direction + moves) % 4
			}
		case 'L':
			{
				moves := val / 90
				moves = moves % 4
				moves = 4 - moves
				direction = (direction + moves) % 4
			}
		default:
			x, y = handleDirection(int(cmd), val, x, y)
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func RunInstructionsWithWayPoint(content string) int {

	instructions := strings.Split(content, "\n")

	dx, dy := 10, 1
	x, y := 0, 0

	changeDirection := func(moves int) {
		for i := 0; i < moves; i++ {
			temp := dx
			dx = dy
			dy = -temp
		}
	}

	for _, instr := range instructions {
		cmd := instr[0]
		val, err := strconv.Atoi(instr[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch cmd {
		case 'F':
			{
				x += dx * val
				y += dy * val
			}
		case 'R':
			{
				moves := val / 90
				moves = moves % 4
				changeDirection(moves)
			}
		case 'L':
			{
				moves := val / 90
				moves = moves % 4
				moves = 4 - moves
				changeDirection(moves)
			}
		default:
			dx, dy = handleDirection(int(cmd), val, dx, dy)
		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func main() {
	content, err := ioutil.ReadFile("./input12.txt")
	if err != nil {
		log.Fatal(err)
	}

	ans := RunInstructions(string(content))

	fmt.Println("Day 12 solution a: ", ans)
	ans = RunInstructionsWithWayPoint(string(content))
	fmt.Println("Day 12 solution b: ", ans)

}
