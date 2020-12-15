package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func FindIterationNumber(content string, final int) int {
	nums := strings.Split(content, ",")
	type Pair struct {
		a,b int
	}
	seen := make(map[int]struct{a, b int})

	last := -1
	lastExists := true
	for index, v := range nums {
		i := index + 1
		vn, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		_, lastExists = seen[vn]
		seen[vn] = Pair{i, -1}
		last = vn
	}

	add := func(v, i int) {
		_, lastExists = seen[v]
		existing, found := seen[v]
		if found {
			existing.b = existing.a
			existing.a = i
			seen[v] = existing
		} else {
			seen[v] = Pair{i, -1}
		}
		//fmt.Println("Add", v, seen[v])
	}


	for index := len(nums); index < final; index++ {
		i := index  + 1
		vn := 0

		if !lastExists {
			add(0, i)
		} else {
			sl := seen[last]
			vn = seen[last].a - seen[last].b
			if sl.b == -1 {
				vn = i - seen[last].a
			}
			add(vn, i)
		}
		last = vn
	}

	return last

}
func main() {
	content := "6,13,1,15,2,0"
	ans := FindIterationNumber(content, 2020)
	fmt.Println(ans)

	ans = FindIterationNumber(content, 30000000)
	fmt.Println(ans)

}
