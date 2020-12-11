package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type OccupiedFn func(grid [][]uint8, ci, cj int) int

func occupiedNeighbors(grid [][]uint8, ci, cj int) int {
	occupied := 0
	cols := len(grid[0])
	rows := len(grid)

	for i := 0 ; i < 3; i++ {
		for j := 0 ; j < 3; j++ {
			ri := ci + i - 1
			rj := cj + j - 1

			if ri == ci && rj == cj {
				continue
			}
			if ri < 0 || ri >= rows {
				continue
			}
			if rj < 0 || rj >= cols {
				continue
			}
			if grid[ri][rj] == '#' {
				occupied++
			}

		}
	}
	return occupied
}

func occupiedRay(grid [][]uint8, ci, cj int, di, dj int) int {
	cols := len(grid[0])
	rows := len(grid)
	ci += di
	cj += dj
	for {
		if ci < 0 || ci >= rows {
			break
		}
		if cj < 0 || cj >= cols {
			break
		}
		if grid[ci][cj] == '#' {
			return 1
		}
		if grid[ci][cj] == 'L' {
			return 0
		}
		ci += di
		cj += dj
	}
	return 0
}

func occupiedNeighborsB(grid [][]uint8, ci, cj int) int {
	occupied := 0

	for i := 0 ; i < 3; i++ {
		for j := 0 ; j < 3; j++ {
			di := i - 1
			dj := j - 1

			if di == 0 && dj == 0 {
				continue
			}
			occupied += occupiedRay(grid, ci, cj, di, dj)
		}
	}
	return occupied
}

func getGrid(content string) [][]uint8 {
	content = strings.TrimSpace(content)
	grid := strings.Split(content, "\n")

	rows := len(grid)

	ans := make([][]uint8, rows)

	for i, line := range grid {
		for _, c := range line {
			ans[i] = append(ans[i], uint8(c))
		}
	}
	return ans

}

func GetSeatsAtFixPoint(content string, tolerance int, fn OccupiedFn) int {
	grid1 := getGrid(content)
	grid2 := getGrid(content)

	cols := len(grid1[0])
	rows := len(grid2)

	changed := true
	currentGrid := &grid1
	nextGrid := &grid2

	for changed {
		changed = false

		temp := currentGrid
		currentGrid = nextGrid
		nextGrid = temp

		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if (*currentGrid)[i][j] == '.' {
					continue
				}
				occupied := fn(*currentGrid, i , j)

				switch (*currentGrid)[i][j] {
				case 'L': {
					if occupied == 0 {
						(*nextGrid)[i][j] = '#'
						changed = true
					} else {
						(*nextGrid)[i][j] = 'L'
					}
				}
				case '#' : {
					if occupied >= tolerance {
						(*nextGrid)[i][j] = 'L'
						changed = true
					} else {
						(*nextGrid)[i][j] = '#'
					}

				}
				}
			}
		}

		if !changed {
			break
		}

	}

	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (*currentGrid)[i][j] =='#' {
				ans++
			}
		}
	}

	return ans
}

func main() {
	content, err := ioutil.ReadFile("./input11.txt")
	if err != nil {
		log.Fatal(err)
	}

	ans := GetSeatsAtFixPoint(string(content), 4, occupiedNeighbors)
	fmt.Println("Solution day 11 part a ", ans)

	ans = GetSeatsAtFixPoint(string(content), 5, occupiedNeighborsB)
	fmt.Println("Solution day 11 part a ", ans)

}
