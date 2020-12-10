package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func getNumbers(content string) []int {
	vec := strings.Split(content, "\n")
	numbers := make([]int, 0)
	for _, v := range vec {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}

	numbers = append(numbers, 0)
	sort.Ints(numbers)
	numbers = append(numbers, numbers[(len(numbers) - 1)] + 3)
	return numbers
}


func UseAllAdapters(content string) int {

	numbers := getNumbers(content)
	threeJoltDif := 0
	oneJoltDif := 0

	for i := 1; i < len(numbers); i++ {

		switch numbers[i] - numbers[i - 1] {
		case 1: {
			println()
			oneJoltDif += 1
		}
		case 2: {
		}
		case 3: {
			threeJoltDif += 1
		}
		default:
			log.Fatalf("Assumption failed, sorted trick doesnt work")
		}


	}
	return threeJoltDif * oneJoltDif
}
func aux (table *map[int]int, graph map[int][]int, cur int) int {

	if cur == 0 {
		return 1
	}
	_, e := graph[cur]; if !e {
		return 0
	}

	v, exists := (*table)[cur]
	if exists {
		return v
	}

	ans := aux(table, graph, cur - 1) + aux(table, graph, cur -2 ) + aux(table, graph, cur -3)

	(*table)[cur] = ans
	return ans
}

func CountPaths(content string) int {
	numbers := getNumbers(content)
	fmt.Println(numbers)
	graph := make(map[int][]int)

	final := numbers[len(numbers) - 1]
	fmt.Println(final)
	for i, a := range numbers {
		for j:= 0; j < i; j++ {
			b := numbers[j]
			if a - b <= 3 && a - b > 0{
				graph[a] = append(graph[a], b)
			}
		}
	}

	table := make(map[int]int)
	ways := aux(&table, graph, final)
	fmt.Println(graph)

	return ways
}

func main() {
	content, err := ioutil.ReadFile("./input10.txt")
	if err != nil {
		log.Fatal(err)
	}
	ans := UseAllAdapters(string(content))
	fmt.Println("Soultion for day 10, a ", ans)

	ans = CountPaths(string(content))
	fmt.Println("Soultion for day 10, b ", ans)

}
