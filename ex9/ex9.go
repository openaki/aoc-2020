package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
	"io/ioutil"
	"log"
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
	return numbers
}

func FindInvalid(content string, offset int) int {
	numbers := getNumbers(content)
	numberSet := make(map[int]int)
	for _, n:= range numbers[:offset] {
		numberSet[n]++
	}

	for i := offset; i < len(numbers); i++ {
		n := numbers[i]
		found := false
		for _, a := range numbers[i-offset:i] {
			_, exists := numberSet[n - a]
			if exists {
				found = true
				break
			}
		}
		if !found {
			return n
		}

		toDelete :=  numbers[i - offset]
		numberSet[toDelete]--
		if numberSet[toDelete] == 0 {
			delete(numberSet, toDelete)
		}
		numberSet[n]++
	}

	return -1
}

func FindKey(content string, key int) int {
	numbers := getNumbers(content)
	i, j := 0, 1
	sum := numbers[i] + numbers[j]
	for {
		if i >= len(numbers) || j >= len(numbers) {
			return -1
		}
		if sum == key {
			fmt.Println(numbers[i:j+1])
			min := intsets.MaxInt
			max := 0
			for _, n := range  numbers[i:j+1] {
				if max < n {
					max = n
				}
				if min > n {
					min = n
				}
			}
			return max + min

		}
		if sum < key {
			j++
			sum += numbers[j]
		}

		if sum > key {
			sum -= numbers[i]
			i++
		}
	}
	return 0
}
func main() {
	content, err := ioutil.ReadFile("./input9.txt")
	if err != nil {
		log.Fatal(err)
	}
	ans := FindInvalid(string(content), 25)
	fmt.Println("Solution day 9 part a ", ans)

	ans = FindKey(string(content), ans)
	fmt.Println("Solution day 9 part a ", ans)

}
