package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Point  struct {
	x, y, z, w int
}

func addPadding(points *map[Point]bool, dim int) map[Point]bool {

	newPoints := *points

	for ep := range *points {
		for x := -1; x < 2; x++ {
			for y := -1; y < 2; y++ {
				for z := -1; z < 2; z++ {
					for w := -1; w < 2; w++ {
						rw := w
						if dim < 4 {
							rw = 0
						}
						if x == 0 && y == 0 && z == 0 {
							continue
						}
						np := Point{ep.x + x, ep.y + y, ep.z + z, ep.w + rw}
						_, exists := newPoints[np]
						if !exists {
							newPoints[np] = false
						}
					}

				}
			}
		}

	}
	return newPoints
}
// .#.
// ..#
// ###
func RunIterations(content string, dim int) int {
	lines := strings.Split(content, "\n")


	type Range struct {
		min, max int
	}

	alive := make(map[Point]bool)
	//xRange := Range {-1, len(lines[0])}
	//yRange := Range {-1, len(lines)}
	//zRange := Range {-1, 1}

	for i, l := range lines {
		for j, c := range l {
			if c == '#' {
				alive[Point{i, j, 0, 0}] = true
			} else {
				alive[Point{i, j, 0, 0}] = false

			}
		}
	}

	alive = addPadding(&alive, dim)
	//fmt.Println(alive)

	iterationCount := 6

	for i := 0; i < iterationCount; i++ {
		nextAlive := make(map[Point]bool)

		for ep, _ := range alive {
			aliveCount := 0
			deadCount := 0

			seen := make(map[Point]bool)

			for x := -1; x < 2; x++ {
				for y := -1; y < 2; y++ {
					for z := -1; z < 2; z++ {
						for w := -1; w < 2; w++ {
							rw := w
							if dim < 4 {
								rw = 0
							}
							if x == 0 && y == 0 && z == 0  && rw == 0{
								continue
							}
							np := Point{ep.x + x, ep.y + y, ep.z + z, ep.w + rw}
							living, exits := alive[np]
							if exits && living && !seen[np]{
								aliveCount++
							} else {
								deadCount++
							}
							if dim < 4 {
								seen[np] = true
							}
						}
					}
				}
			}
			//fmt.Println(ep, aliveCount, deadCount, alives)

			if alive[ep] && (aliveCount == 2 || aliveCount == 3) {
				nextAlive[ep] = true
			} else if aliveCount == 3 {
				nextAlive[ep] = true
			} else {
				nextAlive[ep] = false
			}

		}

		alive = nextAlive
		//fmt.Println(alive)
		alive = addPadding(&alive, dim)


		//for k,v :=range alive {
		//	if v {
		//		fmt.Println(k)
		//	}
		//}
		//break
	}
	count := 0
	for _,v :=range alive {
		if v {
			count++
		}
	}

	return count
}


func main() {
	content, err := ioutil.ReadFile("./input17.txt")
	if err != nil {
		log.Fatal(err)
	}
	ans := RunIterations(string(content), 3)
	fmt.Println(ans)
	ans = RunIterations(string(content), 4)
	fmt.Println(ans)

}
